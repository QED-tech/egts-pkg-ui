package egts

import "github.com/qed-tech/egts-pkg-debugger/internal/usecase"

type Handler struct {
	decodeUsecase     *usecase.DecodeUsecase
	getHistoryUsecase *usecase.GetHistoryUsecase
}

func NewHandler(
	decodeUsecase *usecase.DecodeUsecase,
	getHistoryUsecase *usecase.GetHistoryUsecase,
) *Handler {
	return &Handler{
		decodeUsecase:     decodeUsecase,
		getHistoryUsecase: getHistoryUsecase,
	}
}
