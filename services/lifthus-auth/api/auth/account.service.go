package auth

import (
	"context"
	"lifthus-auth/common/db"
)

// DeleteAccountService takes uid and deletes corresponding Lifthus user account.
func (as authApiService) DeleteAccountService(c context.Context, uid uint64) error {
	err := db.DeleteUser(c, uid)
	if err != nil {
		return err
	}
	return nil
}
