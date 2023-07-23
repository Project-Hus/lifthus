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

// ValidateSessionV2 only validates the session token.
// and returns the session ID, user ID, whether it is expired, and error.
func ValidateSessionV2(ctx context.Context, lst string) (
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

	uidUint, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		return nil, nil, false, fmt.Errorf("parsing uid failed:%w", err)
	}
	uid = &uidUint

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
		return nil, fmt.Errorf("invalid session")
	}
	// get and parse the session ID and TID.
	sidStr := claims["sid"].(string)
	sid, err1 := uuid.Parse(sidStr)
	tidStr := claims["tid"].(string)
	tid, err2 := uuid.Parse(tidStr)
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("invalid session")
	}

	// check if the session is valid by querying the database.
	// and get the user entity too.
	ls, err = db.Client.Session.Query().Where(session.ID(sid)).WithUser().Only(ctx) // WithUser always.
	if err != nil {
		return nil, fmt.Errorf("invalid session")
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
func CreateSessionV2(ctx context.Context) (ls *ent.Session, newSignedToken string, err error) {
	// create new lifthus session
	ns, err := db.Client.Session.Create().Save(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("creating session failed:%w", err)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pps": "lifthus_session",
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
	// if a user is signed to Cloudhus sessino,
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
		case ls.UID != nil && *ls.UID == cuDto.Uid:
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

	var uidStr string
	if cu != nil {
		uidStr = strconv.FormatUint(cu.ID, 10)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pps": "lifthus_session",
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

// ========================================================================================

// CreateLifthusSession creates new session for user and returns the session ID.
func CreateSession(ctx context.Context, client *ent.Client) (sid string, stSigned string, err error) {
	// create new lifthus session
	ns, err := client.Session.Create().Save(ctx)
	if err != nil {
		return "", "", fmt.Errorf("!!creating new session failed:%w", err)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"purpose": "lifthus_session",
		"sid":     ns.ID.String(),
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	hsk := []byte(lifthus.HusSecretKey)
	stSigned, err = st.SignedString(hsk)
	if err != nil {
		return "", "", fmt.Errorf("!!signing session token failed:%w", err)
	}

	return ns.ID.String(), stSigned, nil
}

func RevokeSession(ctx context.Context, client *ent.Client, sid string) error {
	sid_uuid, err := uuid.Parse(sid)
	if err != nil {
		return fmt.Errorf("!!parsing uuid failed:%w", err)
	}
	// delete the session from database
	err = client.Session.DeleteOneID(sid_uuid).Exec(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("!!revoking session failed:%w", err)
	}
	return nil
}

func SignSession(ctx context.Context, client *ent.Client, sid string, uid uint64) error {
	sid_uuid, err := uuid.Parse(sid)
	if err != nil {
		return fmt.Errorf("!!parsing uuid failed: %w", err)
	}
	// update the session with uid
	_, err = client.Session.UpdateOneID(sid_uuid).SetUID(uid). /*SetUsed(false).*/ Save(ctx)
	if err != nil {
		return fmt.Errorf("!!updating session with uid failed: %w", err)
	}
	return nil
}

// UUID version
// func SignSession(ctx context.Context, client *ent.Client, sid string, uid string) error {
// 	sid_uuid, err := uuid.Parse(sid)
// 	uid_uuid, err1 := uuid.Parse(uid)
// 	if err != nil || err1 != nil {
// 		return fmt.Errorf("!!parsing uuid failed: %w, %w", err, err1)
// 	}
// 	// update the session with uid
// 	_, err = client.Session.UpdateOneID(sid_uuid).SetUID(uid_uuid).SetUsed(false).Save(ctx)
// 	if err != nil {
// 		return fmt.Errorf("!!updating session with uid failed: %w", err)
// 	}
// 	return nil
// }

// ValidateSession validates the session and updates the db.
// cases:
// B-1: if it is signed but expired, reset used, signed_at, uid from db and return same SID and empty UID
// B-2: if it is not signed and expired, return same SID
// C: if it is valid, just return
func ValidateSession(ctx context.Context, client *ent.Client, st string) (
	sid string,
	uid string,
	exp bool,
	err error,
) {
	// parse session token
	stParsed, exp, err := helper.ParseJWTWithHMAC(st)
	if err != nil {
		return "", "", false, fmt.Errorf("parsing jwt token failed:%w", err)
	}
	pps, ok := stParsed["purpose"].(string)
	if !ok || pps != "lifthus_session" {
		return "", "", false, fmt.Errorf("parsing jwt token failed: wrong purpose")
	}
	// get sid and uid, if not found, return error
	sid, ok = stParsed["sid"].(string)
	if !ok {
		return "", "", false, fmt.Errorf("parsing jwt token failed: sid not found")
	}
	uid, ok = stParsed["uid"].(string)
	if !ok {
		return "", "", false, fmt.Errorf("parsing jwt token failed: uid not found")
	}
	sid_uuid, err := uuid.Parse(sid)
	if err != nil {
		return "", "", false, fmt.Errorf("parsing uuid failed:%w", err)
	}

	// if it is expired and signed, update the session in db to be unsigned.
	if exp && uid != "" {
		err = client.Session.UpdateOneID(sid_uuid). /*SetUsed(false).*/ ClearSignedAt().ClearUID().Exec(ctx)
		if err != nil {
			return "", "", false, fmt.Errorf("updating session failed:%w", err)
		}
	}
	return sid, uid, exp, nil
}

// RefreshSession refreshes the session token with same SID and empty UID.
func RefreshSessionToken(ctx context.Context, client *ent.Client, sid string) (stSigned string, err error) {
	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"purpose": "lifthus_session",
		"sid":     sid,
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	})
	stSigned, err = st.SignedString([]byte(lifthus.HusSecretKey))
	if err != nil {
		return "", err
	}
	return stSigned, nil
}

// RevokeHusToken takes session token and revokes them.
func RevokeSessionToken(ctx context.Context, client *ent.Client, st string) error {
	stClaims, _, err := helper.ParseJWTWithHMAC(st)
	if err != nil {
		return err
	}

	sid_uuid, err := uuid.Parse(stClaims["sid"].(string))
	if err != nil {
		return err
	}

	err = client.Session.DeleteOneID(sid_uuid).Exec(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	return nil
}
