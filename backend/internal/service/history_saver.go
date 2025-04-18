package service

import (
	"context"
	"fmt"

	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
	"github.com/qed-tech/egts-pkg-debugger/internal/repository"
)

type HistorySaver interface {
	Save(ctx context.Context, hexPackage string) error
}

type historySaver struct {
	historyRepository repository.HistoryRepository
}

func NewHistorySaver(historyRepository repository.HistoryRepository) HistorySaver {
	return &historySaver{
		historyRepository: historyRepository,
	}
}

func (h *historySaver) Save(ctx context.Context, hexPackage string) error {
	err := h.historyRepository.WriteHistory(ctx, &entity.HistoryRecord{
		EgtsPackageHex: hexPackage,
	})

	if err != nil {
		return fmt.Errorf("failed to write history")
	}

	return nil
}
