include .env

APP_NAME=users-app

build-local:
	GOOS=${PLATFORM} GOARCH=amd64 CGO_ENABLED=0 go build -o ${APP_NAME} ./cmd/api

start: build-local
	./${APP_NAME}
