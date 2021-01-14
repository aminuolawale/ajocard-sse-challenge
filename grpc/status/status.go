package status

import (
	"aminuolawale/ajocard/helpers"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetStatus(ctx context.Context, in *Status) (*Status, error) {
	response := helpers.ReadStatus()
	return &Status{Active: response.Active, FailureRate: float32(response.FailureRate), FailureThreshold:float32(response.FailureThreshold)}, nil
}
