# Build
FROM golang:1.14-alpine as builder

WORKDIR /go/src/github.com/farukkaradeniz/go-training

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Run
FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /app/go-training

COPY --from=builder /go/src/github.com/farukkaradeniz/go-training .

CMD ["./main"]