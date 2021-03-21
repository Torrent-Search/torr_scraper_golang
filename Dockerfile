FROM tejasvp25/golang-alpine-ytdl:latest

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/scraper/
RUN mkdir downloads && chmod +x downloads
COPY . .

RUN go mod download
RUN go build main.go