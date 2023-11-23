package main

import (
	"context"
	"life_game/internal/application"
	"os"
)

func main() {
	ctx := context.Background()
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  100,
		Height: 100,
	}
	app := application.New(cfg)
	return app.Run(ctx)
}
