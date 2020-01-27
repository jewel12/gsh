.PHONY: clean

all: gsh

gsh: clean
	go build -o $@

clean:
	-rm gsh