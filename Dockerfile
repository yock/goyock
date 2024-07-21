# syntax=docker/dockerfile:1

FROM golang:bookworm

WORKDIR /app
COPY . ./
RUN go build -o server.go .
CMD ["./server.go"]
