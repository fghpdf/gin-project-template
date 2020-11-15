package connectivity

import (
	"context"
	"fghpdf.me/gin-project-template/internal/pkg/common"
	"fghpdf.me/gin-project-template/internal/pkg/ctxType"
	log "github.com/sirupsen/logrus"
)

type Model struct {
	Status string `json:"status"`
}

type Server interface {
	Ping(ctx context.Context) (*Model, error)
}

type server struct {
}

func NewServer() Server {
	return &server{}
}

// Ping check server status
func (s *server) Ping(ctx context.Context) (*Model, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	model := &Model{Status: "pong"}
	log.WithField(common.REQUEST_ID, ctxType.GetRequestId(ctx)).
		Infof("[connectivity][Ping] params: , response: %v", model)
	return model, nil
}
