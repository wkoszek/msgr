GO=go
all: msgr

msgr: main.go config.go context.go slack.go tele.go helper.go
	$(GO) build -o $@ $^
	
lint:
	golint
