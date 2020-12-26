.PHONY: dep
dep:
	@go get -v -x && go mod download && go mod tidy

.PHONY: fmt
fmt:
	@go fmt

.PHONY: build
build:
	@make dep && GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build

.PHONY: run
run:
	@echo "" && CGO_ENABLED=0 GOARCH=amd64 go run *.go
