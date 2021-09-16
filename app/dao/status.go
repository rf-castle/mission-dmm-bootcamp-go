package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"
)

type (
	// Implementation for repository.Status
	status struct {
		db          *sqlx.DB
		accountRepo repository.Account
	}
)

func NewStatus(db *sqlx.DB, accountRepo repository.Account) repository.Status {
	return &status{db: db, accountRepo: accountRepo}
}

func (r *status) FindById(ctx context.Context, id object.StatusId) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(
		ctx,
		"SELECT * from status where id = ?",
		id,
	).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	user, err := r.accountRepo.FindById(ctx, entity.AccountId)
	if err != nil {
		return nil, err
	}
	entity.Account = user
	return entity, nil
}

func (r *status) Create(ctx context.Context, accountId object.AccountID, content string, mediaIDs ...object.MediaID) (*object.Status, error) {
	var statusId object.StatusId
	result, err := r.db.ExecContext(
		ctx,
		"INSERT INTO status (account_id, content) values (?, ?);",
		accountId, content,
	)
	if err != nil {
		return nil, err
	}
	statusId, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.FindById(ctx, statusId)
}
