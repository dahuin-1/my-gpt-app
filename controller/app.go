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

type PostMessageResult struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     bool   `json:"logprobs,omitempty"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
