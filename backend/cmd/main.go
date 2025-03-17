package main

import (
	"github.com/qed-tech/egts-pkg-debugger/internal/di"
	"github.com/qed-tech/egts-pkg-debugger/pkg/application"
)

func main() {
	container := di.NewContainer()

	app := application.NewApplication(container)
	app.Run()
}
