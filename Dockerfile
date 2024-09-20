FROM golang:1.21

WORKDIR /go/src/github.com/fastenhealth/fasten-sources
COPY . .

RUN go mod tidy && go mod vendor && go build -o /usr/bin/test-smart-client ./tools/test-smart-client/main.go

CMD ["/usr/bin/test-smart-client"]

EXPOSE 9999
