package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("app", func() {
	Title("simple chat app")
	Server("app", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			//URI("grpc://localhost:8080")
		})
	})
})

//var Question = Type("Question", func() {
//	Attribute("topics", String)
//	Required("topics")
//})

var _ = Service("app", func() {
	Description("service charged based on token usage.")
	Error("InternalServerError", ErrorType)
	Error("BadRequest", ErrorType)
	Error("Unauthorized", ErrorType)
	HTTP(func() {
		Response("InternalServerError", StatusInternalServerError)
		Response("BadRequest", StatusBadRequest)
		Response("Unauthorized", StatusUnauthorized)
	})
	cors.Origin("*", func() { // ISSUE: 도메인 결정시 수정
		cors.Headers("*")
		cors.Methods("GET", "POST")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})

	Method("postMessage", func() {
		Security(APIKeyAuth)
		Payload(PostMessagePayload)
		Result(func() {
			Attribute("answer", String)
			Required("answer")
		})
		HTTP(func() {
			POST("/message")
			Header("key:X-API-Key")
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

/*
curl https://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [
      {
        "role": "system",
        "content": "You are a helpful assistant."
      },
      {
        "role": "user",
        "content": "Hello!"
      }
    ]
  }'
*/

var PostMessagePayload = Type("PostMessagePayload", func() {
	APIKey("api_key", "key", String)
	Attribute("model", String)
	Attribute("messages", ArrayOf(String))
	Required("model", "messages")
})

/*
{
  "id": "chatcmpl-123",
  "object": "chat.completion",
  "created": 1677652288,
  "model": "gpt-3.5-turbo-0125",
  "system_fingerprint": "fp_44709d6fcb",
  "choices": [{
    "index": 0,
    "message": {
      "role": "assistant",
      "content": "\n\nHello there, how may I assist you today?",
    },
    "logprobs": null,
    "finish_reason": "stop"
  }],
  "usage": {
    "prompt_tokens": 9,
    "completion_tokens": 12,
    "total_tokens": 21
  }
}
*/

var PostMessageResult = Type("PostMessageResult", func() {
	Attribute("email")
	Attribute("authority")
	Required("email", "authority")
})
