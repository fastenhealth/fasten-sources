.PHONY: deps
deps:
	go mod tidy && go mod vendor

.PHONY: serve-backend
serve-backend: deps
	cd testutils && go run oauth_cli.go

build-js:
	tygo generate

publish-js: build-js
	cd js && tsc
	cd js && npm run pub

test:
	go test ./...
