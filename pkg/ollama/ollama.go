package ollama

import (
	"context"
	"fmt"
	"log/slog"
	"math"

	"github.com/jmorganca/ollama/api"
)

type Ollama struct {
	ModelName string
	*api.Client
}

func New(ctx context.Context, log *slog.Logger, modelName string) (*Ollama, error) {
	c, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	if err := c.Pull(ctx, &api.PullRequest{
		Name: modelName,
	}, progress(log)); err != nil {
		return nil, fmt.Errorf("failed to pull: %w", err)
	}

	log.Info("ollama model pulled: ok")

	return &Ollama{
		ModelName: modelName,
		Client:    c,
	}, nil
}

func (o *Ollama) Generate(ctx context.Context, req *api.GenerateRequest, fn api.GenerateResponseFunc) error {
	req.Model = o.ModelName
	return o.Client.Generate(ctx, req, fn)
}

func progress(log *slog.Logger) func(pr api.ProgressResponse) error {
	return func(pr api.ProgressResponse) error {
		if pr.Total == 0 {
			log.Info(fmt.Sprintf("(?%%) %s", pr.Status))
			return nil
		}
		pc := int(math.Floor(100 * (float64(pr.Completed) / float64(pr.Total))))
		log.Info(fmt.Sprintf("(%d%%) %s", pc, pr.Status))
		return nil
	}
}
