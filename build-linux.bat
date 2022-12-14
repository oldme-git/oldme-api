@echo off
SET GO_ARCH=amd64
SET GOOS=linux
SET CGO_ENABLED=0
go build -o oldme main.go
@pause