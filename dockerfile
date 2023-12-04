FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-api

EXPOSE 6000

CMD ["/go-api"]
