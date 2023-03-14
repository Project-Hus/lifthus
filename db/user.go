package db

import (
	"context"
	"lifthus-auth/ent"
	"lifthus-auth/ent/user"
	"log"

	"github.com/google/uuid"
)

func GetUserByUID(client *ent.Client, uid string) (*ent.User, error) {
	uid_uuid, err := uuid.Parse(uid)
	if err != nil {
		log.Println("[F] parsing uuid failed: ", err)
	}
	u, err := client.User.Query().Where(user.ID(uid_uuid)).Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		log.Println("[F] getting user by uid failed: ", err)
		return nil, err
	}
	return u, nil
}
