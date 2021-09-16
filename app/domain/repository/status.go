package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	FindById(context.Context, object.StatusId) (*object.Status, error)
	Create(context.Context, object.AccountID, string, ...object.MediaID) (*object.Status, error)
	Delete(context.Context, object.StatusId) error
}
