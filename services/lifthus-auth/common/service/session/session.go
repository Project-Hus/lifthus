package session

import (
	"context"
	"fmt"
	"lifthus-auth/common/db"
	"lifthus-auth/common/helper"
	"lifthus-auth/common/lifthus"
	"lifthus-auth/ent"
	"lifthus-auth/ent/session"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func IsExpired(err error) bool {
	return err != nil && err.Error() == "expired session"
}

// ValidateSession gets Lifthus session token in string and validates it.
// if token is invalid, it returns "invalid session" error.
// if token is expired but vaild except the expiration issue, it returns "expired session" error with session entity. (IsExpired func is provided to check it)
// if revoked token is used, it returns "illegal session" error.
// and if it is valid, it returns Lifthus session with User edge.
func ValidateSessionV2(ctx context.Context, lst string) (ls *ent.Session, err error) {
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

	// if session is valid regardless of expiration, return expiration error with session entity to try refreshing the session.
	if exp {
		return ls, fmt.Errorf("expired sesison")
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
		"purpose":   "lifthus_session",
		"sid":       ns.ID.String(),
		"tid":       ns.Tid.String(),
		"connected": false, // it tells that it is still not connected to Hus session.
		"uid":       "",
		"exp":       time.Now().Add(time.Minute * 5).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	hsk := []byte(lifthus.HusSecretKey)
	stSigned, err := st.SignedString(hsk)
	if err != nil {
		return nil, "", fmt.Errorf("signing session token failed:%w", err)
	}

	return ns, stSigned, nil
}

// RefreshSessionHard gets old Lifthus session and refreshes it.
// it queries the DB to verify whether the user is still signed and etc.
// the term Hard means that it does not only check Lifthus DB but it also double checks Cloudhus DB to verify whether the user is still signed.
func RefreshSessionHard(ctx context.Context, ols *ent.Session) (nls *ent.Session, newSignedToken string, err error) {
	lu := ols.QueryUser().OnlyX(ctx)
	if lu != nil {
		// do api call to cloudhus to check the session
	}
	// change tid, and issue new token etc
	return nil, "", nil
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
		"uid":     "",
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
	_, err = client.Session.UpdateOneID(sid_uuid).SetUID(uid).SetUsed(false).Save(ctx)
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
		err = client.Session.UpdateOneID(sid_uuid).SetUsed(false).ClearSignedAt().ClearUID().Exec(ctx)
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
		"uid":     "",
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
