package postgres

import (
	"context"
	"fmt"

	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{
		db: db,
	}
}

func (r HistoryRepository) Conn() *gorm.DB {
	return r.db
}

func (r HistoryRepository) GetHistory(ctx context.Context, db *gorm.DB, limit int) ([]*entity.HistoryRecord, error) {
	var (
		out []*entity.HistoryRecord
	)

	err := r.Conn().
		WithContext(ctx).
		Table("egts_history_records").
		Order("id DESC").
		Limit(limit).
		Scan(&out).
		Error

	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	return out, nil
}

func (r HistoryRepository) WriteHistory(ctx context.Context, historyRecord *entity.HistoryRecord) error {
	err := r.Conn().
		WithContext(ctx).
		Save(historyRecord).
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Debug().
		Error

	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}

	return nil
}
