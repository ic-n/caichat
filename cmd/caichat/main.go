package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ic-n/caichat/pkg/api"
	"github.com/ic-n/caichat/pkg/ollama"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	errLog := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	ollamaClient, err := ollama.New(ctx, log, os.Getenv("OLLAMA_MODEL"))
	if err != nil {
		errLog.Error("failed to create ollama client: %s", err)
		return
	}

	h, err := api.NewHandler(ctx, api.NewServer(log, ollamaClient))
	if err != nil {
		errLog.Error("failed to create server handler: %s", err)
		return
	}

	s := http.Server{
		Addr:              ":80",
		Handler:           h,
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 5,
	}
	if err := s.ListenAndServe(); err != nil {
		errLog.Error("failed to serve: %s", err)
	}
}
