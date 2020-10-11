FROM golang:1.15 as builder

WORKDIR /app
COPY go.mod go.sum ./
COPY vendor ./vendor

COPY internal ./internal
COPY main.go ./

RUN go test -race ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags="-w -s" main.go

FROM gcr.io/distroless/base:nonroot
COPY --from=builder \
  /go/bin/main \
  /usr/local/bin/feed
ENTRYPOINT ["/usr/local/bin/feed"]
