FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go run github.com/steebchen/prisma-client-go generate
RUN go mod tidy