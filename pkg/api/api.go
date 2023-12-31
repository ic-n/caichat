package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	caichatv1 "github.com/ic-n/caichat/pkg/api/gen/caichat/v1"
	"github.com/ic-n/caichat/pkg/ollama"
	"github.com/jmorganca/ollama/api"
)

type Server struct {
	log          *slog.Logger
	ollamaClient *ollama.Ollama
	caichatv1.UnimplementedCaiChatServiceServer
}

func (s Server) Health(_ context.Context, _ *caichatv1.HealthRequest) (*caichatv1.HealthResponse, error) {
	return &caichatv1.HealthResponse{
		Ok: true,
	}, nil
}

func (s Server) Generate(ctx context.Context, req *caichatv1.GenerateRequest) (*caichatv1.GenerateResponse, error) {
	var r strings.Builder

	s.log.Info("got request")

	err := s.ollamaClient.Generate(ctx, &api.GenerateRequest{
		Prompt: req.Prompt,
		System: "You are exceptionally short, you use only few words to reply.",
	}, func(gr api.GenerateResponse) error {
		s.log.With("gr", gr).Info("generated")
		r.WriteString(gr.Response)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &caichatv1.GenerateResponse{
		Text: r.String(),
	}, nil
}

func NewServer(log *slog.Logger, ollamaClient *ollama.Ollama) *Server {
	return &Server{log: log, ollamaClient: ollamaClient}
}

func NewHandler(ctx context.Context, s *Server) (http.Handler, error) {
	mux := runtime.NewServeMux()
	if err := caichatv1.RegisterCaiChatServiceHandlerServer(ctx, mux, s); err != nil {
		return nil, fmt.Errorf("failed to register service: %w", err)
	}

	return mux, nil
}
