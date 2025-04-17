package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/qed-tech/egts-pkg-debugger/internal/di"
	"github.com/qed-tech/egts-pkg-debugger/pkg/application"
)

func main() {
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal("failed to init container")
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	app := application.NewApplication(container)
	app.Run()
}
