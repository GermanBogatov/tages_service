package main

import (
	"context"
	"fmt"
	"github.com/GermanBogatov/tages_service/app/internal/app"
	"github.com/GermanBogatov/tages_service/app/internal/config"
	"github.com/GermanBogatov/tages_service/app/pkg/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.Info(ctx, "config initializing")
	cfg := config.GetConfig()

	fmt.Println(cfg)
	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		logging.Fatal(ctx, err)
	}

	logging.Info(ctx, "Running Application")
	a.Run(ctx)
}
