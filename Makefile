CUR_DIR=$(shell pwd)
CMD_DIR=cmd/$(cmd)
BIN_DIR_LINUX=bin/$(cmd)/linux
BIN_DIR_WIN=bin/$(cmd)/win
BIN_DIR_DARWIN=bin/$(cmd)/darwin

GO_BUILD=go build

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BIN_DIR_LINUX)/$(cmd) -v $(CMD_DIR)/main.go
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BIN_DIR_WIN)/$(cmd).exe -v $(CMD_DIR)/main.go
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o $(BIN_DIR_DARWIN)/$(cmd) -v $(CMD_DIR)/main.go