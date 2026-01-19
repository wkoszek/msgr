GO := go
BINARY := msgr

.PHONY: all build build-linux clean fmt vet test

all: fmt vet build

build:
	$(GO) build -o $(BINARY) ./cmd/msgr

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY).linux ./cmd/msgr

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test -race ./...

clean:
	rm -f $(BINARY) $(BINARY).linux
