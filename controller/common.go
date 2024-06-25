package controller

import (
	"context"
	"log"
	common "my-gpt/gen/common"
)

// common service example implementation.
// The example methods log the requests and return zero values.
type commonsrvc struct {
	logger *log.Logger
}

// NewCommon returns the common service implementation.
func NewCommon(logger *log.Logger) common.Service {
	return &commonsrvc{logger}
}

// Health implements health.
func (s *commonsrvc) Health(ctx context.Context) (err error) {
	s.logger.Print("common.health")
	return
}
