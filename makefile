GO=go
all: msgr

msgr: main.go config.go slack.go tele.go helper.go
	$(GO) build -o $@ $^
	
