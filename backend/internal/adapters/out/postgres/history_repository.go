package postgres

import (
	"context"

	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
	"gorm.io/gorm"
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
	return []*entity.HistoryRecord{
		{
			ID:  1,
			Hex: "0100020b002300020001871800010011e70300000202101500b6739d1b4fba3a9ed227bc350000000000000000003b07",
		},
		{
			ID:  2,
			Hex: "0100000B004800C0AD01423900143485860500004F5FE51002021015000F5EE51020181B9EB4DFDA341920825F0000000001110A001F0700060000001303001B0700010800C30600001B070002090003000000D82A",
		},
	}, nil
}
