## build stage
#FROM golang:alpine as builder
#
#WORKDIR /build
#
#COPY . .
#
#RUN go mod download
#
#RUN go build -o ./main ./cmd/app/main.go

# product stage
FROM golang:alpine

WORKDIR /depeche-backend

COPY . .

RUN apk update

RUN go mod download

RUN apk add bash
RUN apk add make
RUN apk add build-base

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

EXPOSE 8081
