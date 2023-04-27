FROM golang:alpine AS builder
WORKDIR /depeche-backend
COPY  . .
RUN apk update
RUN CGO_ENABLED=0 GOOS=linux go build -o /main -ldflags="-s -w" -a -installsuffix cgo ./cmd/app/main.go

FROM alpine
COPY --from=builder /main /

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "-c" , "mv /main /depeche-backend && /depeche-backend/main"]