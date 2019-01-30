#!/usr/bin/env sh

echo "`go version`"

echo "making builds... linux"
GOOS=linux GOARCH=amd64 go build capcoins-api.go
mv capcoins-api capcoins-api_linux

echo "making builds... mac"
GOOS=darwin GOARCH=amd64 go build capcoins-api.go
mv capcoins-api capcoins-api_mac

echo "making builds... windows"
GOOS=windows GOARCH=amd64 go build capcoins-api.go
mv capcoins-api.exe capcoins-api_win.exe

echo "making builds... -> default"
go build capcoins-api.go

