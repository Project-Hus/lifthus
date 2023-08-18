package auth

import (
	"context"
	"fmt"
	"lifthus-auth/common/db"
	"lifthus-auth/common/helper"
	"lifthus-auth/common/lifthus"
	"lifthus-auth/common/service/session"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (as authApiService) signOutService(c context.Context, lst string) (nlst string, err error) {
	ls, err := session.ValidateSessionQueryUser(c, lst)
	if session.IsExpiredValid(err) {
		return "", fmt.Errorf("session expird: %w", err)
	} else if err != nil {
		return "", fmt.Errorf("failed to validate session: %w", err)
	} else if ls.Edges.User == nil {
		return "", fmt.Errorf("the session is not signed")
	}

	propagCh := make(chan error) // for waiting propagation goroutine
	txCh := make(chan error)     // for waiting transaction goroutine

	go func() {
		sot, err := helper.SignedHusTotalSignOutToken(ls.Hsid.String())
		if err != nil {
			propagCh <- fmt.Errorf("failed to generate token")
			return
		}
		// request to the Cloudhus endpoint
		req, err := http.NewRequest(http.MethodPatch, "https://auth.cloudhus.com/auth/hus/signout"+"", strings.NewReader(sot))
		if err != nil {
			propagCh <- fmt.Errorf("failed to create request")
			return
		}
		resp, err := lifthus.Http.Do(req)
		if err != nil {
			propagCh <- fmt.Errorf("propagation failed")
			return
		}
		defer resp.Body.Close()
		// propagation failed or succeeded
		if resp.StatusCode == http.StatusOK {
			propagCh <- nil
			return
		}
		propagCh <- fmt.Errorf("propagation failed")
	}()

	tx, err := db.Client.Tx(c)
	if err != nil {
		return "", fmt.Errorf("failed to start transaction")
	}

	go func() {
		ls, err = tx.Session.UpdateOne(ls).ClearUID().ClearSignedAt().SetTid(uuid.New()).Save(c)
		if err != nil {
			db.Rollback(tx, err)
			txCh <- err
			return
		}
		txCh <- nil
	}()

	propagErr := <-propagCh
	txErr := <-txCh

	if propagErr != nil || txErr != nil {
		_ = db.Rollback(tx, nil)
		return "", fmt.Errorf("failed to sign out propag or tx failed")
	}

	err = tx.Commit()
	if err != nil {
		_ = db.Rollback(tx, err)
		return "", fmt.Errorf("failed to commit tx")
	}

	lst, err = session.SessionToToken(c, ls)
	if err != nil {
		return "", fmt.Errorf("failed to tokenize the session")
	}

	return lst, nil
}
