package usecase

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/kuznetsovin/egts-protocol/libs/egts"
	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
)

type DecodeUsecase struct {
}

func NewDecodeUsecase() *DecodeUsecase {
	return &DecodeUsecase{}
}

var (
	ErrInvalidHexString = errors.New("invalid hex string")
)

func (u *DecodeUsecase) Decode(pkg string) (*entity.DebugPackage, error) {
	bytes, err := hex.DecodeString(pkg)
	if err != nil {
		return nil, ErrInvalidHexString
	}

	egtsPackage := egts.Package{}
	_, err = egtsPackage.Decode(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode bytes to egts package, err: %v", err)
	}

	return entity.NewDebugPackage(&egtsPackage), nil
}
