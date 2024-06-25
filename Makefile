APP_NAME_API = my-gpt-app
OUT_DIR = ./out

build-api: goagen
	go build -o $(OUT_DIR)/$(APP_NAME_API) ./cmd/api

goagen:
	rm -rf my-gpt/gen
	goa gen my-gpt/design