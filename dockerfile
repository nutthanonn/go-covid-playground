FROM golang:1.18

WORKDIR /go/src/go-gin-api

COPY . .

RUN go build -o bin/server cmd/app.go

CMD ["./bin/server"]