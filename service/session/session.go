package session

import (
	"context"
	"fmt"
	"lifthus-auth/ent"
	"lifthus-auth/helper"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// CreateLifthusSession creates new session for user and returns the session ID.
func CreateSession(ctx context.Context, client *ent.Client) (sid string, stSigned string, err error) {
	// create new lifthus session
	ns, err := client.Session.Create().Save(ctx)
	if err != nil {
		return "", "", fmt.Errorf("!!creating new session failed:%w", err)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": ns.ID.String(),
		"uid": "",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	hsk := []byte(os.Getenv("HUS_SECRET_KEY"))
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

func SetSignedSession(ctx context.Context, client *ent.Client, sid string, uid string) error {
	sid_uuid, err := uuid.Parse(sid)
	uid_uuid, err1 := uuid.Parse(uid)
	if err != nil || err1 != nil {
		return fmt.Errorf("!!parsing uuid failed: %w, %w", err, err1)
	}
	// update the session with uid
	_, err = client.Session.UpdateOneID(sid_uuid).SetUID(uid_uuid).Save(ctx)
	if err != nil {
		return fmt.Errorf("!!updating session with uid failed: %w", err)
	}
	return nil
}

// ValidateSessionToken validates the session and updates the db.
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
	stParsed, exp, err := helper.ParseJWTwithHMAC(st)
	if err != nil {
		return "", "", false, fmt.Errorf("!!parsing jwt token failed:%w", err)
	}
	// get sid and uid, if not found, return error
	sid, ok := stParsed["sid"].(string)
	if !ok {
		return "", "", false, fmt.Errorf("!!parsing jwt token failed: sid not found")
	}
	uid, ok = stParsed["uid"].(string)
	if !ok {
		return "", "", false, fmt.Errorf("!!parsing jwt token failed: uid not found")
	}
	sid_uuid, err := uuid.Parse(sid)
	if err != nil {
		return "", "", false, fmt.Errorf("!!parsing uuid failed:%w", err)
	}

	// if it is expired and signed, update the session in db to be unsigned.
	if exp && uid != "" {
		err = client.Session.UpdateOneID(sid_uuid).SetUsed(false).ClearSignedAt().ClearUID().Exec(ctx)
		if err != nil {
			return "", "", false, fmt.Errorf("!!updating session failed:%w", err)
		}
	}
	return sid, uid, exp, nil
}

// RefreshSession refreshes the session token with same SID and empty UID.
func RefreshSessionToken(ctx context.Context, client *ent.Client, sid string) (stSigned string, err error) {
	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": sid,
		"uid": "",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	stSigned, err = st.SignedString([]byte(os.Getenv("HUS_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return stSigned, nil
}

// RevokeHusToken takes session token and revokes them.
func RevokeSessionToken(ctx context.Context, client *ent.Client, st string) error {
	stClaims, _, err := helper.ParseJWTwithHMAC(st)

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
