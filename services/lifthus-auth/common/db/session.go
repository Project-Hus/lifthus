package db

import (
	"context"
	"fmt"
	"lifthus-auth/ent"
	"lifthus-auth/ent/session"

	"github.com/google/uuid"
)

// QuerySessionByID queries the session by session's ID.
// if an error occurs, it returns nil, error.
// if session is not found, it returns nil, nil.
// if session is found, it returns session, nil.
func QuerySessionByID(c context.Context, sid uuid.UUID) (*ent.Session, error) {
	ls, err := Client.Session.Query().Where(session.ID(sid)).Only(c)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("querying session failed:%w", err)
	}
	return ls, nil
}

// QuerySessionByHsid queries the latest session that is related to hsid.
// if an error occurs, it returns nil, error.
// if session is not found, it returns nil, nil.
// if session is found, it returns session, nil.
//
// Hus session may last longer so multiple sessions may be related to one hsid if there is no proper disconnection process.
// so this function returns the latest one.
func QuerySessionByHsid(c context.Context, hsid uuid.UUID) (*ent.Session, error) {
	// get latest session related to hsid
	lss, err := Client.Session.Query().Where(session.Hsid(hsid)).Order(ent.Desc(session.FieldConnectedAt)).First(c)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("querying session failed:%w", err)
	}
	return lss, nil
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
