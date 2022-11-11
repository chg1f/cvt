.PHONY: build
build:
	go build -o cvt ./main.go

.PHONY: test
test:
	go test -v ./...

