# Dockerfile
# Dockerfile
FROM golang:latest

RUN apt-get update && apt-get install -y wget gnupg curl

RUN curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb

RUN apt-get install -y ./google-chrome-stable_current_amd64.deb

RUN rm google-chrome-stable_current_amd64.deb 

WORKDIR /app

COPY . .

RUN go install github.com/sensepost/gowitness@latest

RUN go build -o main .

CMD ["./main"]

# Step 1: Install debian with chrome
# FROM debian:bullseye-slim


# RUN curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb

# RUN apt-get install -y ./google-chrome-stable_current_amd64.deb

# RUN rm google-chrome-stable_current_amd64.deb 

# # Step 2: Install Golang

# ENV GOROOT /usr/local/go

# ENV GOPATH /go

# ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# RUN wget -qO- https://go.dev/dl/go1.21.5.linux-amd64.tar.gz | tar -C /usr/local -xz

# # Create a workspace directory for your Go projects
# RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# # Set the working directory
# WORKDIR $GOPATH/src

# RUN go version

# # Step 3: Run app
# WORKDIR /app

# COPY . .

# RUN go install github.com/sensepost/gowitness@latest

# RUN go build -o main

# CMD ["./main"]
