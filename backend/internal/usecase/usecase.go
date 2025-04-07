package usecase

import "github.com/qed-tech/egts-pkg-debugger/internal/service"

type Usecase struct {
	DecodeUsecase *DecodeUsecase
}

func NewUsecase() *Usecase {
	return &Usecase{
		DecodeUsecase: NewDecodeUsecase(
			service.NewHistorySaver(),
		),
	}
}
