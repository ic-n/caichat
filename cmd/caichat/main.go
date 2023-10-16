package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ic-n/caichat/pkg/api"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	errLog := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	h, err := api.NewHandler(ctx, api.NewServer(log))
	if err != nil {
		errLog.Error("failed to create server handler: %s", err)
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
