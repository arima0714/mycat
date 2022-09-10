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
	go test -shuffle=on -count=1
