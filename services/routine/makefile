.DEFAULT_GOAL := build

go: # run the app with nodemon for hot reload
	HUS_ENV=native nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run ./cmd/routine/main.go
.PHONY:go

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
	# shadow ./... # this tool detects shadowing variables
.PHONY:vet

build: vet
	GOWORK=off go generate ./internal/ent
	go build ./cmd/routine/...
.PHONY:build