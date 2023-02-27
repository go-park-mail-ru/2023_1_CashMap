FROM golang:alpine AS builder
WORKDIR /depeche-backend
COPY . .
RUN apk update
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w" -a -installsuffix cgo ./cmd/app/main.go

FROM alpine
WORKDIR /depeche-backend
COPY --from=builder ./depeche-backend .
# TODO: понять, через какой порт подключиться к контейнер

ENTRYPOINT ["./main"]
