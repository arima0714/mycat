.PHONY: all
all: build

.PHONY: build
build:
	go build

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test -v -shuffle=on -count=1
