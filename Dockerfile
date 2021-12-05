# Dockerfile distroless
FROM golang:1.14 as builder
WORKDIR /go/src/main
COPY . .
RUN go build -ldflags="-s -w" -o main main.go && cp main /go/bin/main


# Building a tiny image
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/main /
CMD ["/main"]

