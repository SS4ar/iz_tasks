FROM golang:1.21.0 as builder

WORKDIR /app

COPY main.go .
COPY go.mod .
COPY go.sum .
COPY resources ./resources
COPY templates ./templates
COPY git ./git

RUN go build -o app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app ./
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/resources ./resources
COPY --from=builder /app/git ./git

RUN apk add libc6-compat

CMD ["./app"]
