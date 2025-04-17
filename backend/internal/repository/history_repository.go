package repository

import (
	"context"

	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	Conn() *gorm.DB
	GetHistory(ctx context.Context, db *gorm.DB, limit int) ([]*entity.HistoryRecord, error)
}
