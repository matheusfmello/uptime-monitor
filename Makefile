APP_NAME=uptime-monitor
CMD_DIR=cmd/server

.PHONY: run build tidy fmt test

run:
	go run ./$(CMD_DIR)

build:
	go build -o bin/$(APP_NAME) ./$(CMD_DIR)

fmt:
	gofmt -w .

tidy:
	go mod tidy

test:
	go test -v ./...

clean:
	rm -rf bin/
