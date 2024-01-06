# Dockerfile
FROM golang:latest

RUN apt-get install -y wget

RUN wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb

RUN apt-get install -y ./google-chrome-stable_current_amd64.deb

WORKDIR /app

COPY . .

RUN go install github.com/sensepost/gowitness@latest

RUN go build -o main

CMD ["./main"]

