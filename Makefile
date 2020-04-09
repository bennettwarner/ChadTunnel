# ref: https://vic.demuzere.be/articles/golang-makefile-crosscompile/

BIN_DIR := build

default: clean darwin linux windows integrity

clean:
	$(RM) $(BIN_DIR)/chadtunnel*
	go clean -x

install:
	go install

darwin:
	GOOS=darwin GOARCH=amd64 go build -o '$(BIN_DIR)/chadtunnel-darwin-amd64'

linux:
	GOOS=linux GOARCH=amd64 go build -o '$(BIN_DIR)/chadtunnel-linux-amd64'

windows:
	GOOS=windows GOARCH=amd64 go build -o '$(BIN_DIR)/chadtunnel-windows-amd64.exe'

integrity:
	cd $(BIN_DIR) && shasum *