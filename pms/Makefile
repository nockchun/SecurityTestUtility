BUILD=go build -ldflags="-w -s"

default: build

build:
	@echo "Building PMS Server ..."
	GOOS=linux GOARCH=amd64 $(BUILD) pms_server.go -o bin/pms_server_linux
	GOOS=windows GOARCH=amd64 $(BUILD) pms_server.go -o bin/pms_server_windows

clean:
	@rm -rf bin/