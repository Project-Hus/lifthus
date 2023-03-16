package session

import (
	"context"
	"fmt"
	"lifthus-auth/ent"
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

	sid = ns.ID.String()
	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": sid,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
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
		return fmt.Errorf("!!deleting session failed:%w", err)
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
