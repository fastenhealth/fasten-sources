.PHONY: deps
deps:
	go mod tidy && go mod vendor

.PHONY: serve-backend
serve-backend: deps
	cd testutils && go run oauth_cli.go
test:
	go test ./...

test-js:
	cd js && yarn run e2e

# Steps related to building and publishing fasten-sources-js library.

build-js:
	go install github.com/gzuidhof/tygo@latest
	tygo generate
	npm install -g typescript
	cd js && yarn install
	cd js && tsc
