package http

import "github.com/qed-tech/egts-pkg-debugger/internal/usecase"

type Handler struct {
	usecases usecase.Usecase
}

func NewHandler() *Handler {
	return &Handler{}
}
