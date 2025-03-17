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
	Usecases *usecase.Usecase
}

func NewContainer() *Container {
	return &Container{
		Adapters: &Adapters{
			HTTP: http.NewHandler(),
		},
		Usecases: usecase.NewUsecase(),
	}
}
