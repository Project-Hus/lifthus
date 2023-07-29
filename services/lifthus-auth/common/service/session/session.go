package session

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"lifthus-auth/common/db"
	"lifthus-auth/common/dto"
	"lifthus-auth/common/helper"
	"lifthus-auth/common/lifthus"
	"lifthus-auth/ent"
	"lifthus-auth/ent/session"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Session Error represents the error that occurs in session service package.
type SessionError struct {
	Message string
}

func (e SessionError) Error() string {
	return e.Message
}

// ExpiredSessionError occurs when the session token is expired.
var ExpiredValidSessionError = &SessionError{"expired valid session"}

// IsExpiredValid checks if the error is ExpiredValidSessionError.
func IsExpiredValid(err error) bool {
	return err == ExpiredValidSessionError
}

// ValidateSession only validates the session token.
// and returns the session ID, user ID, whether it is expired, and error.
func ValidateSession(ctx context.Context, lst string) (
	sid *uuid.UUID,
	uid *uint64,
	exp bool,
	err error,
) {
	// parse session token
	stParsed, exp, err := helper.ParseJWTWithHMAC(lst)
	if err != nil {
		return nil, nil, false, fmt.Errorf("parsing jwt token failed:%w", err)
	}
	pps, ok := stParsed["pps"].(string)
	if !ok || pps != "lifthus_session" {
		return nil, nil, false, fmt.Errorf("parsing jwt token failed: wrong purpose")
	}
	// get sid and uid, if not found, return error
	sidStr, ok := stParsed["sid"].(string)
	if !ok {
		return nil, nil, false, fmt.Errorf("parsing jwt token failed: sid not found")
	}
	uidStr, ok := stParsed["uid"].(string)
	if !ok {
		return nil, nil, false, fmt.Errorf("parsing jwt token failed: uid not found")
	}
	suuid, err := uuid.Parse(sidStr)
	sid = &suuid
	if err != nil {
		return nil, nil, false, fmt.Errorf("parsing uuid failed:%w", err)
	}

	if uidStr != "" {
		uidUint, err := strconv.ParseUint(uidStr, 10, 64)
		if err != nil {
			return nil, nil, false, fmt.Errorf("parsing uid failed:%w", err)
		}
		uid = &uidUint
	}

	return sid, uid, exp, nil
}

// ValidateSessionQueryUser gets Lifthus session token in string and validates it.
// if token is invalid, it returns "invalid session" error.
// if token is expired but vaild except the expiration issue, it returns "expired valid session" error with session entity. (IsExpiredValid func is provided to check it)
// if revoked token is used, it returns "illegal session" error.
// and if it is valid, it returns Lifthus session with User edge.
func ValidateSessionQueryUser(ctx context.Context, lst string) (ls *ent.Session, err error) {
	// parse the Lifthus session token.
	claims, exp, err := helper.ParseJWTWithHMAC(lst)
	if err != nil || claims["pps"].(string) != "lifthus_session" {
		return nil, fmt.Errorf("invalid session: %w", err)
	}
	// get and parse the session ID and TID.
	sidStr := claims["sid"].(string)
	sid, err1 := uuid.Parse(sidStr)
	tidStr := claims["tid"].(string)
	tid, err2 := uuid.Parse(tidStr)
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("invalid session: %w", err)
	}

	// check if the session is valid by querying the database.
	// and get the user entity too.
	ls, err = db.Client.Session.Query().Where(session.ID(sid)).WithUser().Only(ctx) // WithUser always.
	if err != nil {
		return nil, fmt.Errorf("invalid session:%w", err)
	}

	if tid != ls.Tid {
		// revoke all user's session and propagate (not implemented yet) ------------------------------------------------------------------------
		return nil, fmt.Errorf("illegal session")
	}

	// if session is valid regardless of expiration, return EV error with session entity to try refreshing the session.
	if exp {
		return ls, ExpiredValidSessionError
	}

	return ls, nil
}

// CreateSession issues new Lifthus session and returns the session entity and signed session token.
func CreateSession(ctx context.Context) (ls *ent.Session, newSignedToken string, err error) {
	// create new lifthus session
	ns, err := db.Client.Session.Create().Save(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("creating session failed:%w", err)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pps": "lifthus_session",
		"iss": "https://auth.lifthus.com",
		"sid": ns.ID.String(),
		"tid": ns.Tid.String(),
		"uid": "",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	stSigned, err := st.SignedString(lifthus.HusSecretKeyBytes)
	if err != nil {
		return nil, "", fmt.Errorf("signing session token failed:%w", err)
	}

	return ns, stSigned, nil
}

// SessionToToken takes lifthus session and returns signed session token.
func SessionToToken(ctx context.Context, ls *ent.Session) (lst string, err error) {
	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pps": "lifthus_session",
		"iss": "https://auth.lifthus.com",
		"sid": ls.ID.String(),
		"tid": ls.Tid.String(),
		"uid": helper.UIDToString(ls.Edges.User),
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	stSigned, err := st.SignedString(lifthus.HusSecretKeyBytes)
	if err != nil {
		return "", fmt.Errorf("signing session token failed:%w", err)
	}

	return stSigned, nil
}

// RefreshSession gets Lifthus session and refreshes it.
// Unlike RefreshSessionHard, it doesn't query Cloudhus API but only Lifthus DB.
func RefreshSession(ctx context.Context, ls *ent.Session) (nls *ent.Session, newSignedToken string, err error) {
	return nil, "", nil
}

// RefreshSessionHard gets Lifthus session and refreshes it.
// it queries the DB to verify whether the user is still signed and etc.
// the term Hard means that it does not only check Lifthus DB but it also double checks Cloudhus API to verify whether the user is signed.
//
// if the user is signed, the user entity must be included in the edges.
//
// if the user turns out not to be registered, it does user-registration process as well.
func RefreshSessionHard(ctx context.Context, ls *ent.Session) (nls *ent.Session, newSignedToken string, err error) {
	// genreate session connection token
	sct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pps":     "hus_connection",
		"iss":     "https://auth.lifthus.com",
		"service": "lifthus",
		"sid":     ls.ID.String(),
		"exp":     time.Now().Add(time.Minute * 10).Unix(),
	})
	sctSigned, err := sct.SignedString(lifthus.HusSecretKeyBytes)
	if err != nil {
		return nil, "", fmt.Errorf("signing session connection token failed:%w", err)
	}

	// from Cloudhus endpoint get the connected session information
	req, err := http.NewRequest(http.MethodGet, "https://auth.cloudhus.com/auth/hus/connect/"+sctSigned, nil)
	if err != nil {
		return nil, "", fmt.Errorf("creating new request failed:%w", err)
	}
	resp, err := lifthus.Http.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("hus connection api failed:%w", err)
	}
	defer resp.Body.Close()
	// if code not 200, return invalid session error
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("invalid session")
	}

	// decode the JSON response
	var husConn dto.HusConnDto
	err = json.NewDecoder(resp.Body).Decode(&husConn)
	if err != nil {
		return nil, "", fmt.Errorf("decoding hus connection response failed:%w", err)
	}

	// parse connected Hus session ID
	hsid, err := uuid.Parse(husConn.Hsid)
	if err != nil {
		return nil, "", fmt.Errorf("invalid hsid")
	}
	// connected user (nil if not signed)
	cuDto := husConn.User

	/* transaction */
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		err = db.Rollback(tx, err)
		return nil, "", fmt.Errorf("starting transaction failed:%w", err)
	}

	trx := tx.Session.UpdateOne(ls).SetHsid(hsid).SetTid(uuid.New())
	var cu *ent.User
	// if the user is newly signed, update it.
	// cases: cu != nil basically
	// ls.Uid == nil -> new user signed
	// ls.Uid != nil, ls.Uid == cu.Uid -> maintain session
	// ls.Uid != nil, ls.Uid != cu.Uid -> update session user
	//
	// if a user is signed to Cloudhus session,
	if cuDto != nil {
		// query the user
		cu, err = db.QueryUserByID(ctx, cuDto.Uid)
		if err != nil {
			err = db.Rollback(tx, err)
			return nil, "", fmt.Errorf("querying user failed:%w", err)
		}
		// query succeeded but user not found, register the user
		if cu == nil {
			cu, err = db.RegisterUser(ctx, *cuDto)
			if err != nil {
				err = db.Rollback(tx, err)
				return nil, "", fmt.Errorf("registering user failed:%w", err)
			}
		}
		switch {
		case ls.UID == nil:
			fallthrough
		case ls.UID != nil && *ls.UID != cuDto.Uid:
			trx = trx.SetUID(cuDto.Uid).SetSignedAt(time.Now())
			//case ls.UID != nil && *ls.UID == cuDto.Uid: // do nothing
		}
	}

	nls, err = trx.Save(ctx)
	if err != nil {
		err = db.Rollback(tx, err)
		return nil, "", fmt.Errorf("refreshing session failed:%w", err)
	}

	nls, err = db.Client.Session.Query().Where(session.ID(nls.ID)).WithUser().Only(ctx)
	if err != nil {
		err = db.Rollback(tx, err)
		return nil, "", fmt.Errorf("querying session failed:%w", err)
	}

	var uidStr string // "" if not signed
	if cu != nil {
		uidStr = strconv.FormatUint(cu.ID, 10)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pps": "lifthus_session",
		"iss": "https://auth.lifthus.com",
		"sid": nls.ID.String(),
		"tid": nls.Tid.String(),
		"uid": uidStr,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	stSigned, err := st.SignedString(lifthus.HusSecretKeyBytes)
	if err != nil {
		err = db.Rollback(tx, err)
		return nil, "", fmt.Errorf("signing session token failed:%w", err)
	}

	err = tx.Commit()
	if err != nil {
		err = db.Rollback(tx, err)
		return nil, "", fmt.Errorf("committing transaction failed:%w", err)
	}

	return nls, stSigned, nil
}
