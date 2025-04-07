package di

import (
	"github.com/qed-tech/egts-pkg-debugger/internal/adapters/http"
	"github.com/qed-tech/egts-pkg-debugger/internal/usecase"
)

type Adapters struct {
	HTTP *http.Handler
}

type Container struct {
	Adapters *Adapters
}

func NewContainer() *Container {
	
	usecases := usecase.NewUsecase()

	return &Container{
		Adapters: &Adapters{
			HTTP: http.NewHandler(usecases),
		},
	}
}
