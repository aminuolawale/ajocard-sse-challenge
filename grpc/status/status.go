package status

import (

	"golang.org/x/net/context"
	"aminuolawale/ajocard/helpers"
)

type Server struct {
}

func (s *Server) GetStatus(ctx context.Context, in *Status) (*Status, error) {
	response := helpers.CheckStatus()
	return &Status{Active: response.Active, FailureRate: float32(response.FailureRate)}, nil
}
