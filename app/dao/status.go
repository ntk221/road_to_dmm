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
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// Create new status
func (r *status) Create(ctx context.Context, status *object.Status) error {
	result, err := r.db.ExecContext(ctx, "insert into status (account_id, content) values (?, ?)", status.AccountID, status.Content)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	status.ID = lastID

	return nil
}

// FindByStatusID : ステータスIDからステータスを取得
func (r *status) FindByStatusID(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
