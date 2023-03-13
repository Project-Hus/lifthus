package session

import (
	"context"
	"fmt"
	"lifthus-auth/ent"
	"log"

	"github.com/google/uuid"
)

// CreateLifthusSession creates new session for user and returns the session ID.
func CreateSession(ctx context.Context, client *ent.Client) (string, error) {
	// create new lifthus session
	ns, err := client.Session.Create().Save(ctx)
	if err != nil {
		log.Println("[F] creating new session failed: ", err)
	}
	return ns.ID.String(), nil
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
