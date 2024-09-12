.PHONY: build run-api run-client
.SILENT:

build:
	go build -o ./build/api ./cmd/api/main.go
	go build -o ./build/client ./cmd/client/main.go

run-api: build
	export JWT_KEY=09aa18a89dee21de8382b9ef169dc5d39130d7036e8b61f20a350ffb0da330cf && \
	./build/api

run-client: build
	./build/client