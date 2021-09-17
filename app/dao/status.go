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

func (r *status) Delete(ctx context.Context, id object.StatusId) error {
	_, err := r.db.ExecContext(
		ctx,
		"DELETE FROM status WHERE id = ?",
		id,
	)
	return err
}

func (r *status) GetPublic(context context.Context, filter *object.TimeLineFilter) ([]*object.Status, error) {
	ret := make([]*object.Status, 0, filter.Limit)
	whereStmt := "WHERE :since_id < id"
	if filter.MaxId >= 0 {
		whereStmt += " && id < :max_id"
	}
	stmt := "SELECT * FROM status " + whereStmt + " ORDER BY id DESC LIMIT :limit"
	rows, err := r.db.NamedQueryContext(context, stmt, filter)
	if err != nil {
		return nil, err
	}
	var i int
	for i = 0; rows.Next(); i++ {
		entity := new(object.Status)
		err = rows.StructScan(entity)
		if err != nil {
			return nil, err
		}
		ret = append(ret, entity)
	}
	return ret, nil
}
