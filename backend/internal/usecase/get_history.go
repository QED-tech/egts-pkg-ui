package usecase

import (
	"context"
	"fmt"

	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
	"github.com/qed-tech/egts-pkg-debugger/internal/repository"
)

type GetHistoryUsecase struct {
	historyRepository repository.HistoryRepository
}

func NewGetHistoryUsecase(historyRepository repository.HistoryRepository) *GetHistoryUsecase {
	return &GetHistoryUsecase{
		historyRepository: historyRepository,
	}
}

const (
	defaultHistoryLimit = 10
)

func (u GetHistoryUsecase) Handle(ctx context.Context) ([]*entity.HistoryRecord, error) {
	history, err := u.historyRepository.GetHistory(
		ctx,
		u.historyRepository.Conn(),
		defaultHistoryLimit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get history from db, err: %v", err)
	}

	return history, nil
}
