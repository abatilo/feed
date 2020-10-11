FROM golang:1.15-alpine as builder

WORKDIR /app
COPY go.mod go.sum vendor ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags="-w -s" main.go

FROM gcr.io/distroless/base:nonroot
COPY --from=builder \
  /go/bin/main \
  /usr/local/bin/feed
ENTRYPOINT ["/usr/local/bin/feed"]
