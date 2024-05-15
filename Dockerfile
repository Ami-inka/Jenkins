FROM golang:1.18-alpine as builder

WORKDIR /app

COPY server.go .

RUN go build -o webserver server.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/webserver .

EXPOSE 1234

CMD ["./webserver"]

# Test commit - mail