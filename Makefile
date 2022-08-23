.PHONY: build clean deploy

install:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o /usr/local/bin/typing typing/main.go
