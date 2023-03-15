package db

import (
	"context"
	"lifthus-auth/common"
	"lifthus-auth/ent"
	"lifthus-auth/ent/user"
	"log"
	"time"

	"github.com/google/uuid"
)

func CreateNewLifthusUser(c context.Context, client *ent.Client, nu common.HusSessionCheckBody) (*ent.User, error) {
	uid_uuid, err := uuid.Parse(nu.Uid)
	if err != nil {
		log.Println("[F] parsing uuid failed: ", err)
		return nil, err
	}
	// create new lifthus user
	lu := client.User.Create().SetID(uid_uuid).
		SetEmail(nu.Email).
		SetEmailVerified(nu.EmailVerified).
		SetName(nu.Name).
		SetGivenName(nu.GivenName).
		SetFamilyName(nu.FamilyName)
	if nu.Birthdate != "" {
		t, err := time.Parse(time.RFC3339, nu.Birthdate)
		if err != nil {
			log.Println("[F] parsing birthdate failed: ", err)
			return nil, err
		}
		lu.SetBirthdate(t)
	}
	nlu, err := lu.Save(c)
	if err != nil {
		log.Println("[F] creating new lifthus userfailed: ", err)
		return nil, err
	}
	return nlu, nil
}

func QueryUserByUID(c context.Context, client *ent.Client, uid string) (*ent.User, error) {
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
