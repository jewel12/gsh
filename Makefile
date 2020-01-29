.PHONY: clean test

all: gsh

gsh: clean
	go build -o $@

clean:
	-rm gsh

test:
	go test ./...