# Dockerfile

# Step 1: Install debian with chrome
FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install -y wget gnupg curl

RUN curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb

RUN apt-get install -y ./google-chrome-stable_current_amd64.deb

RUN rm google-chrome-stable_current_amd64.deb 

# Step 2: Install Golang

RUN wget -qO- https://go.dev/dl/go1.21.5.linux-amd64.tar.gz | tar -C /usr/local -xz

ENV GOROOT /usr/local/go

ENV GOPATH /go

ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# Step 3: Run app
WORKDIR /app

COPY . .

RUN go install github.com/sensepost/gowitness@latest

RUN go build -o main

CMD ["./main"]
