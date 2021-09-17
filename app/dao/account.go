package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.
		QueryRowxContext(ctx, "select * from account where username = ?", username).
		StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *account) FindById(ctx context.Context, userId object.AccountID) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.
		QueryRowxContext(ctx, "select * from account where id = ?", userId).
		StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// Create : ユーザを作成
func (r *account) Create(ctx context.Context, newAccount *object.Account) (*object.Account, error) {
	result, err := r.db.NamedExecContext(
		ctx,
		"insert into account (username, password_hash) values (:username, :password_hash)",
		newAccount,
	)
	if err != nil {
		return nil, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.FindById(ctx, userId)
}
