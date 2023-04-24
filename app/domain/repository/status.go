package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Create new status
	Create(ctx context.Context, status *object.Status) error
	FindByStatusID(ctx context.Context, id int64) (*object.Status, error)
}
