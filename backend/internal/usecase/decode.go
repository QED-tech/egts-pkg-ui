package usecase

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"

	"github.com/kuznetsovin/egts-protocol/libs/egts"
	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
	"github.com/qed-tech/egts-pkg-debugger/internal/service"
)

type DecodeUsecase struct {
	historySaver service.HistorySaver
}

func NewDecodeUsecase(historySaver service.HistorySaver) *DecodeUsecase {
	return &DecodeUsecase{
		historySaver: historySaver,
	}
}

var (
	ErrInvalidHexString = errors.New("invalid hex string")
)

func (u *DecodeUsecase) Decode(ctx context.Context, hexString string) (*entity.DebugPackage, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, ErrInvalidHexString
	}

	egtsPackage := egts.Package{}
	_, err = egtsPackage.Decode(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode bytes to egts package, err: %v", err)
	}

	if err := u.historySaver.Save(ctx, hexString); err != nil {
		slog.WarnContext(ctx, "failed to save history", slog.Any("error", err))
	}

	return entity.NewDebugPackage(&egtsPackage, hexString), nil
}
