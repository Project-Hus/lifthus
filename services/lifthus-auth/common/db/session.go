package db

import (
	"context"
	"fmt"
	"lifthus-auth/ent"
	"lifthus-auth/ent/session"

	"github.com/google/uuid"
)

func QuerySessionByID(c context.Context, sid uuid.UUID) (*ent.Session, error) {
	ls, err := Client.Session.Query().Where(session.ID(sid)).Only(c)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("querying session failed:%w", err)
	}
	return ls, nil
}

// ===========

func QuerySessionBySID(c context.Context, client *ent.Client, sid string) (*ent.Session, error) {
	sid_uuid, err := uuid.Parse(sid)
	if err != nil {
		return nil, fmt.Errorf("!!parsing uuid failed:%w", err)
	}
	s, err := client.Session.Query().Where(session.ID(sid_uuid)).Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("!!getting session failed:%w", err)
	}
	return s, nil
}
