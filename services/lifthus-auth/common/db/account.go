package db

import "context"

func DeleteUser(c context.Context, uid uint64) error {
	err := Client.User.DeleteOneID(uid).Exec(c)
	if err != nil {
		return err
	}
	return nil
}
