GO=go
all: msgr

msgr: main.go config.go context.go slack.go tele.go helper.go mailgun.go
	$(GO) build -o $@ $^
	
lint:
	golint
