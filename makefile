GO=go
all: msgr msgr.linux

msgr: main.go config.go context.go slack.go tele.go helper.go mailgun.go twilio.go
	gofmt -w *.go
	$(GO) build -o $@ $^

msgr.linux: main.go config.go context.go slack.go tele.go helper.go mailgun.go twilio.go
	gofmt -w *.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GO) build -o $@ $^

lint:
	golint
