#!/bin/bash

echo "Compiling the code for windows ..."
GOOS=windows GOARCH=386 go build -o GoHttpDebugRequest.exe main.go
echo "Compiling the code for linux ..."
go build -o GoHttpDebugRequest.exe main.go
echo "Compiling the code for mac ..."
GOOS=darwin GOARCH=amd64 go build -o GoHttpDebugRequest.mac
