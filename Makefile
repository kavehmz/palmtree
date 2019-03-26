.PHONY: $(shell ls -d *)

default:
	@echo "Usage: make [command]"

test:
	docker run --rm -u`id -u`:`id -g` -v $$PWD:/go/src/github.com/kavehmz/palmtree golang:1 /bin/bash -c \
	cd /go/src/github.com/kavehmz/palmtree && \
	go test -v --race -cover -coverprofile=cover.out ./... && \
	go tool cover -func=cover.out | \
		awk 'END {sub("[.].*","",$$NF); printf "Coverage: %d%%\n", $$NF; \
			if ($$NF+0 < 100) {print "Coverage is not sufficient"; exit 1}}'

lint:
	docker run --rm -v $$PWD:/go/src/github.com/kavehmz/palmtree -w /go/src/github.com/kavehmz/palmtree \
		golangci/golangci-lint:latest golangci-lint run ./...
