package main

import (
	"log/slog"
	"os"

	"github.com/nokamoto/analyticapr-go/internal/infra/command"
	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
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
	res, err := command.NewGh().ListPulls(&v1.Repository{
		Owner: "nokamoto",
		Repo:  "analyticapr-go",
	}, "2024-03-06")
	if err != nil {
		slog.Error("failed to list pulls", "error", err)
		os.Exit(1)
	}
	slog.Info("pulls", slog.Any("pulls", res))
}
