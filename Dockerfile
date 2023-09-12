FROM golang:1.21

WORKDIR /go/src/github.com/fastenhealth/fasten-sources
COPY . .

RUN go mod tidy && go mod vendor && go build -o /usr/bin/oauth_cli ./testutils/oauth_cli.go

CMD ["/usr/bin/oauth_cli"]

EXPOSE 9999
