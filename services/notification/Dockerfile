FROM golang:1.21.3-alpine3.18


ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=0

RUN mkdir app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD go run .