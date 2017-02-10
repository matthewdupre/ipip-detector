GO_BUILD_CONTAINER?=calico/go-build:v0.2

all:
	mkdir -p bin
	CGO_ENABLED=0 go build -v -o bin/ipip-detector .
	docker build -t ipip-detector:latest .

.PHONY: clean
clean:
	rm -f bin/ipip-detector
