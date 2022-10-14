FROM golang:1.19.2-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY cmd/main/main.go ./

RUN go build -o /go-bananas

EXPOSE 8080

CMD ["/go-bananas"]