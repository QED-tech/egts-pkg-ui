package usecase

type Usecase struct {
	DecodeUsecase *DecodeUsecase
}

func NewUsecase() *Usecase {
	return &Usecase{
		DecodeUsecase: NewDecodeUsecase(),
	}
}
