# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

COPY ./ ./

RUN go build -o main ./api/
RUN ls
ENTRYPOINT [ "./main" ]