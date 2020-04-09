#! /bin/bash

GOOS=darwin GOARCH=amd64 go build -o dist/chadtunnel-mac src/main.go && echo -e "\033[32m [+] Darwin build complete: $(file dist/chadtunnel-mac)\n\033[0m"

GOOS=linux GOARCH=amd64 go build -o dist/chadtunnel-linux src/main.go && echo -e "\033[32m [+] Linux build complete: $(file dist/chadtunnel-linux)\n\033[0m"

GOOS=windows GOARCH=amd64 go build -o dist/chadtunnel-windows.exe src/main.go && echo -e "\033[32m [+] Windows build complete: $(file dist/chadtunnel-windows.exe)\n\033[0m"

GOOS=darwin GOARCH=amd64