package main

import (
	"log/slog"
	"os"

	"github.com/qed-tech/egts-pkg-debugger/internal/di"
	"github.com/qed-tech/egts-pkg-debugger/pkg/application"
)

func main() {
	container := di.NewContainer()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	app := application.NewApplication(container)
	app.Run()
}
