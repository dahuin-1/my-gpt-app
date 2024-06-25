package controller

import (
	"context"
	"fmt"
	"log"
	"my-gpt/gen/app"

	"goa.design/goa/v3/security"
)

// gpt service example implementation.
// The example methods log the requests and return zero values.
type gptsrvc struct {
	logger *log.Logger
}

// NewGpt returns the gpt service implementation.
func NewGpt(logger *log.Logger) app.Service {
	return &gptsrvc{logger}
}

func (s *gptsrvc) makeError(e ErrorType) error {
	return &app.ErrorType{ErrorCode: e.ErrorCode, Message: e.Message, Status: e.Status}
}

// JWTAuth implements the authorization logic for service "gpt" for the "jwt"
// security scheme.
func (s *gptsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return ctx, fmt.Errorf("not implemented")
}

// PostMessage implements postMessage.
func (s *gptsrvc) PostMessage(ctx context.Context, p *app.PostMessagePayload) (res *app.PostMessageResult, err error) {
	res = &app.PostMessageResult{}
	s.logger.Print("gpt.postMessage")
	return
}

type PostMessagePayload struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}
