package session

import (
	"context"
	"lifthus-auth/ent"
	"log"
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
