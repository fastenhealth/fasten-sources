.PHONY: dep-backend
dep-backend:
	go mod vendor

.PHONY: serve-backend
serve-backend: dep-backend
	cd testutils && go run oauth_cli.go
