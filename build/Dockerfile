FROM golang:1.24-bookworm AS build

WORKDIR /app

RUN apt-get update && apt-get install -y git protobuf-compiler build-essential

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY internal/ ./internal

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o ./bin/main ./cmd

FROM debian:bookworm-slim AS app

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 5001
EXPOSE 50051

COPY --from=build /app/bin/main .
COPY config/ ./config

CMD ["./main"]
