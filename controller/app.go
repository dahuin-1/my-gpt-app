package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"goa.design/goa/v3/security"
	"log"
	"my-gpt/gen/app"
	"my-gpt/pkg/db"
	"net/http"
)

var client http.Client

// gpt service example implementation.
// The example methods log the requests and return zero values.
type gptsrvc struct {
	logger *log.Logger
	client http.Client
}

// NewGpt returns the gpt service implementation.
func NewGpt(logger *log.Logger) app.Service {
	return &gptsrvc{logger, client}
}

func (s *gptsrvc) makeError(e ErrorType) error {
	return &app.ErrorType{ErrorCode: e.ErrorCode, Message: e.Message, Status: e.Status}
}

// JWTAuth implements the authorization logic for service "gpt" for the "jwt"
// security scheme.
func (s *gptsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return ctx, fmt.Errorf("not implemented")
}

func (s *gptsrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	apiKey := viper.GetString("auth.api_key")
	if apiKey != key {
		return nil, s.makeError(ErrInvalidApiKey)
	}
	return ctx, nil
}

// PostMessage implements postMessage. https://platform.openai.com/docs/api-reference/chat/create
func (s *gptsrvc) PostMessage(ctx context.Context, p *app.PostMessagePayload) (res *app.PostMessageResult, err error) {
	res = &app.PostMessageResult{}
	s.logger.Print("postMessage")
	//api call
	requestByte, err := json.Marshal(&db.PostMessagePayload{
		Model:    p.Model,
		Messages: p.Messages,
	})
	if err != nil {
		return res, s.makeError(ErrCanNotMarshal)
	}
	req, err := http.NewRequest(http.MethodPost, viper.GetString("url"), bytes.NewBuffer(requestByte))
	if err != nil {
		return res, s.makeError(ErrCanNotRequest)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.Key))
	response, err := s.client.Do(req)
	if err != nil {
		return res, s.makeError(ErrInternal)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return res, s.makeError(ErrInternal)
	}
	//res

	return
}
