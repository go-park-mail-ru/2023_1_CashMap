pFROM golang:alpine AS builder
WORKDIR /depeche-backend
COPY . .

RUN apk update && apk add bash && apk add make
RUN CGO_ENABLED=0 GOOS=linux go build -o /main -ldflags="-s -w" -a -installsuffix cgo ./cmd/app/main.go
RUN rm -rf /depeche-backend/*
RUN mv /main .

EXPOSE 8080

CMD ["tail", "-f","/dev/null"]