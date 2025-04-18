package di

import (
	"github.com/qed-tech/egts-pkg-debugger/internal/adapters/in/http/egts"
	"github.com/qed-tech/egts-pkg-debugger/internal/adapters/out/postgres"
	"github.com/qed-tech/egts-pkg-debugger/internal/repository"
	"github.com/qed-tech/egts-pkg-debugger/internal/service"
	"github.com/qed-tech/egts-pkg-debugger/internal/usecase"
	postgresdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Container struct {
	Adapters *Adapters
}

type Adapters struct {
	EGTS *egts.Handler
}

type Repositories struct {
	HistoryRepository repository.HistoryRepository
}

type Usecase struct {
	DecodeUsecase     *usecase.DecodeUsecase
	GetHistoryUsecase *usecase.GetHistoryUsecase
}

func NewUsecase(repos Repositories) *Usecase {
	return &Usecase{
		DecodeUsecase: usecase.NewDecodeUsecase(
			service.NewHistorySaver(repos.HistoryRepository),
		),
		GetHistoryUsecase: usecase.NewGetHistoryUsecase(repos.HistoryRepository),
	}
}

func NewContainer() (*Container, error) {
	db, err := initDatabase()
	if err != nil {
		return nil, err
	}

	repositories := Repositories{
		HistoryRepository: postgres.NewHistoryRepository(db),
	}

	usecases := NewUsecase(repositories)

	container := &Container{
		Adapters: &Adapters{
			EGTS: egts.NewHandler(
				usecases.DecodeUsecase,
				usecases.GetHistoryUsecase,
			),
		},
	}

	return container, nil
}

func initDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=egts port=5432 sslmode=disable TimeZone=UTC"
	return gorm.Open(postgresdriver.Open(dsn), &gorm.Config{})
}
