FROM golang:alpine AS builder
WORKDIR /depeche-backend
COPY  . .
RUN apk update
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w" -a -installsuffix cgo ./cmd/app/main.go

FROM alpine
WORKDIR /depeche-backend
COPY --from=builder ./depeche-backend/main .
COPY --from=builder ./depeche-backend/.env .

RUN mkdir ./configs

COPY --from=builder ./depeche-backend/configs ./configs

EXPOSE 8080


ENTRYPOINT ["./main"]