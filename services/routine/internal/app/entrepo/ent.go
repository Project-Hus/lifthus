package entrepo

import (
	"context"
	"fmt"
	"log"
	"routine/internal/ent"
	"routine/internal/repository"
	"routine/pkg/db"
)

func NewEntRepository() *EntRepository {
	return &EntRepository{c: db.EntClient()}
}

type EntRepository struct {
	c  *ent.Client
	tx *ent.Tx
}

func (repo *EntRepository) BeginOrContinueTx(ctx context.Context) (finally func(), err error) {
	if repo.tx != nil {
		return repo.txFinallyContinue(), nil
	}
	tx, err := repo.c.Tx(ctx)
	if err != nil {
		return repo.txFinallyContinue(), err
	}
	repo.tx = tx
	return repo.txFinallyRollback(), nil
}

func (repo *EntRepository) Commit() error {
	if repo.tx == nil {
		return repository.ErrNoTransaction
	}
	err := repo.tx.Commit()
	if err != nil {
		return err
	}
	repo.tx = nil
	return nil
}

func (repo *EntRepository) Rollback(err error) error {
	if repo.tx == nil {
		return repository.ErrNoTransaction
	}
	if rerr := repo.tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	repo.tx = nil
	return err
}

func (repo *EntRepository) txFinallyRollback() func() {
	return func() {
		if repo.tx == nil {
			return
		}
		err := fmt.Errorf("tx is not committed")
		err = repo.Rollback(err)
		if err != nil {
			log.Printf("failed to rollback tx: %v", err)
		}
	}
}

func (repo *EntRepository) txFinallyContinue() func() {
	return func() {
	}
}
