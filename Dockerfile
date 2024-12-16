FROM golang:1.23.4 AS builder
WORKDIR /workspace/go
COPY ./ /workspace/go

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /build/main ./cmd/main.go

FROM golang:1.23.4 AS local
WORKDIR /workspace/go
COPY --from=builder /workspace/go /workspace/go
RUN go install github.com/air-verse/air@latest

FROM gcr.io/distroless/static@sha256:f05686e02ba3e9ff0d947c5ec4ec9d8f00a4bfae0309a2704650db7dca8d6c48 as prod
COPY --from=builder /build/main /app/main
WORKDIR /app
CMD ["./main"]