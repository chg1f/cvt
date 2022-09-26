.PHONY: build
build:
	go build -o bin/cvt cmd/cvt/main.go

.PHONY: test
test:
	go test -v ./...

