FROM tejasvp25/alpine-golang-alpine-docker:latest
WORKDIR $GOPATH/src/github.com/scraper/
COPY . .
RUN go build main.go