package main

import (
	"log/slog"
	"os"

	"github.com/nokamoto/analyticapr-go/internal/infra/command"
	"github.com/nokamoto/analyticapr-go/internal/infra/config"
	"github.com/nokamoto/analyticapr-go/internal/usecase"
)

func main() {
	var level slog.Level
	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		if err := level.UnmarshalText([]byte(os.Getenv("LOG_LEVEL"))); err != nil {
			panic(err)
		}
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})))
	gh := command.NewGh()
	file := "config.yaml"
	if s := os.Getenv("CONFIG_FILE"); s != "" {
		file = s
	}
	cfg, err := config.NewConfig(file)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	app := usecase.NewAnalyticapr(gh, cfg)
	res, err := app.GetAnalytica()
	if err != nil {
		slog.Error("failed to get analyticapr", "error", err)
		os.Exit(1)
	}
	slog.Info("analyticapr", "result", res)
}
