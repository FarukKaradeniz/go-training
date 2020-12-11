FROM golang:1.14-alpine

WORKDIR /go/src/github.com/farukkaradeniz/go-training

COPY . .

RUN go build -o main

CMD ["./main"]