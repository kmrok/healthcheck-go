FROM golang:latest AS builder
WORKDIR /go/src/healthcheck-go
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM alpine:latest
COPY --from=builder /go/src/healthcheck-go/app ./app
ENTRYPOINT ["./app"]
