package db

import (
	"context"
	"fmt"
	"lifthus-auth/common/dto"
	"lifthus-auth/ent"
	"lifthus-auth/ent/user"
	"strconv"
	"time"
)

func CreateNewLifthusUser(c context.Context, client *ent.Client, nu dto.HusSessionCheckBody) (*ent.User, error) {
	// create new lifthus user
	nuUid, err := strconv.ParseUint(nu.Uid, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("!!parsing uid failed:%w", err)
	}
	lu := client.User.Create().SetID(nuUid).
		SetProfileImageURL(nu.ProfileImageURL).
		SetEmail(nu.Email).
		SetEmailVerified(nu.EmailVerified).
		SetName(nu.Name).
		SetGivenName(nu.GivenName).
		SetFamilyName(nu.FamilyName)
	if nu.Birthdate != "" {
		t, err := time.Parse(time.RFC3339, nu.Birthdate)
		if err != nil {
			return nil, fmt.Errorf("!!parsing birthdate failed:%w", err)
		}
		lu.SetBirthdate(t)
	}
	nlu, err := lu.Save(c)
	if err != nil {
		return nil, fmt.Errorf("!!creating new lifthus userfailed:%w", err)
	}
	return nlu, nil
}

// UUID version
// func CreateNewLifthusUser(c context.Context, client *ent.Client, nu common.HusSessionCheckBody) (*ent.User, error) {
// 	uid_uuid, err := uuid.Parse(nu.Uid)
// 	if err != nil {
// 		return nil, fmt.Errorf("!!parsing uuid failed:%w", err)
// 	}
// 	// create new lifthus user
// 	lu := client.User.Create().SetID(uid_uuid).
// 		SetEmail(nu.Email).
// 		SetEmailVerified(nu.EmailVerified).
// 		SetName(nu.Name).
// 		SetGivenName(nu.GivenName).
// 		SetFamilyName(nu.FamilyName)
// 	if nu.Birthdate != "" {
// 		t, err := time.Parse(time.RFC3339, nu.Birthdate)
// 		if err != nil {
// 			return nil, fmt.Errorf("!!parsing birthdate failed:%w", err)
// 		}
// 		lu.SetBirthdate(t)
// 	}
// 	nlu, err := lu.Save(c)
// 	if err != nil {
// 		return nil, fmt.Errorf("!!creating new lifthus userfailed:%w", err)
// 	}
// 	return nlu, nil
// }

func QueryUserByUID(c context.Context, client *ent.Client, uid uint64) (*ent.User, error) {
	u, err := client.User.Query().Where(user.ID(uid)).Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("!!getting user by uid failed:%w", err)
	}
	return u, nil
}

func QueryUserByUsername(c context.Context, client *ent.Client, username string) (*ent.User, error) {
	u, err := client.User.Query().Where(user.Username(username)).Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("!!getting user by username failed:%w", err)
	}
	return u, nil
}

// UUID version
// func QueryUserByUID(c context.Context, client *ent.Client, uid string) (*ent.User, error) {
// 	uid_uuid, err := uuid.Parse(uid)
// 	if err != nil {
// 		return nil, fmt.Errorf("!!parsing uuid failed:%w", err)
// 	}
// 	u, err := client.User.Query().Where(user.ID(uid_uuid)).Only(context.Background())
// 	if err != nil && !ent.IsNotFound(err) {
// 		return nil, fmt.Errorf("!!getting user by uid failed:%w", err)
// 	}
// 	return u, nil
// }

func UpdateUserInfo(c context.Context, client *ent.Client, userInfo dto.UpdateUserInfoDto) (*ent.User, error) {
	u, err := client.User.UpdateOneID(userInfo.Uid).
		SetNillableUsername(userInfo.Username).
		SetNillableBirthdate(userInfo.Birthdate).
		SetNillableCompany(userInfo.Company).
		SetNillableLocation(userInfo.Location).
		SetNillableContact(userInfo.Contact).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("!!updating user info failed:%w", err)
	}
	return u, nil
}
