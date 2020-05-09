FROM golang:alpine
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/scraper/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/PuerkitoBio/goquery
# Build the binary.
RUN go build main.go