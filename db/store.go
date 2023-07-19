package db

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/binsabit/billing/sqlc/models"
)

type Store struct {
	*models.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: models.New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(queries *models.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := models.New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rbErr: %v \n txErr: %v", rbErr, err)
		}
		return err
	}
	tx.Commit()
	return nil
}
