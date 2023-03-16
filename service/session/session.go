package session

import (
	"context"
	"fmt"
	"lifthus-auth/ent"
	"log"
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
		log.Println("[F] creating new session failed: ", err)
	}
	sid = ns.ID.String()

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": sid,
		"uid": nil, // it will be omitted actually.
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hsk := []byte(os.Getenv("HUS_SECRET_KEY"))
	stSigned, err = st.SignedString(hsk)
	if err != nil {
		log.Println("[F] signing session token failed: ", err)
		return "", "", err
	}

	return ns.ID.String(), stSigned, nil
}

func RevokeSession(ctx context.Context, client *ent.Client, sid string) error {
	sid_uuid, err := uuid.Parse(sid)
	if err != nil {
		err = fmt.Errorf("[F]parsing uuid failed:%w", err)
		log.Println(err)
		return err
	}
	// delete the session from database
	err = client.Session.DeleteOneID(sid_uuid).Exec(ctx)
	if err != nil && !ent.IsNotFound(err) {
		err = fmt.Errorf("[F]deleting session failed:%w", err)
		log.Println(err)
		return err
	}
	return nil
}

func SetSignedSession(ctx context.Context, client *ent.Client, sid string, uid string) error {
	sid_uuid, err := uuid.Parse(sid)
	uid_uuid, err1 := uuid.Parse(uid)
	if err != nil || err1 != nil {
		log.Println("[F] parsing uuid failed: ", err, err1)
		return fmt.Errorf("parsing uuid failed: %w, %w", err, err1)
	}
	// update the session with uid
	_, err = client.Session.UpdateOneID(sid_uuid).SetUID(uid_uuid).Save(ctx)
	if err != nil {
		log.Println("[F] updating session with uid failed: ", err)
		return fmt.Errorf("updating session with uid failed: %w", err)
	}
	return nil
}
