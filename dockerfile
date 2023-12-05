FROM golang:alpine as builder

WORKDIR /app

COPY ../../.. .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder app .

ENTRYPOINT ./app -host=postgres -port=5432 -user=admindb -password=password -dbname=postgres -sslmode=true