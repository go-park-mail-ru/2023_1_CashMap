#!/bin/bash

go mod download

go build -o ./main ./cmd/app/main.go
