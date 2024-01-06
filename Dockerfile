# Dockerfile
FROM golang:latest

WORKDIR /app

COPY . .

RUN go install github.com/sensepost/gowitness@latest

RUN go build -o main

CMD ["./main"]

