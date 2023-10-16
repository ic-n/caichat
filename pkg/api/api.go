package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	caichatv1 "github.com/ic-n/caichat/pkg/api/gen/caichat/v1"
)

type Server struct {
	log *slog.Logger
	caichatv1.UnimplementedCaiChatServiceServer
}

func (s Server) Health(_ context.Context, _ *caichatv1.HealthRequest) (*caichatv1.HealthResponse, error) {
	s.log.With(slog.Attr{
		Key:   "endpoint",
		Value: slog.StringValue("health"),
	}).Info("called")

	return &caichatv1.HealthResponse{
		Ok: true,
	}, nil
}

func NewServer(log *slog.Logger) *Server {
	return &Server{log: log}
}

func NewHandler(ctx context.Context, s *Server) (http.Handler, error) {
	mux := runtime.NewServeMux()
	if err := caichatv1.RegisterCaiChatServiceHandlerServer(ctx, mux, s); err != nil {
		return nil, fmt.Errorf("failed to register service: %w", err)
	}

	return mux, nil
}
