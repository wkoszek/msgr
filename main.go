package main

import "fmt"
import "flag"
import "os"
import "bufio"
import "strings"


func main() {
	var	ctx	Context

	ctx.ArgWhere = flag.String("where", "slack", "service to post [slack]")
	ctx.ArgProfileName = flag.String("profile", "default", "profile name [default]")
	ctx.ArgCode = flag.Bool("code", false, "paste code block [disabled]")
	ctx.ArgHTML = flag.Bool("html", false, "use HTML for emails")
	ctx.ArgFrom = flag.String("from", "default", "From field (email)")
	ctx.ArgTo = flag.String("to", "default", "To field (email)")
	ctx.ArgSubject = flag.String("subject", "default", "Subject field (email)")

	flag.Parse()

	validWheres := []string{ "slack", "telegram", "mailgun" }
	found := false
	for _, v := range validWheres {
		if *ctx.ArgWhere == v {
			found = true
		}
	}
	if !found {
		fmt.Printf("Option '%s' unsupported, supported types are: %s\n",
			*ctx.ArgWhere, strings.Join(validWheres, ", "));
		os.Exit(1)
	}

	println("before config")

	cfg := NewConfig(".msgr.conf", *ctx.ArgProfileName)
	ctx.Config = cfg

	println("before getMsg")
	msgStr := getMsg()

	if *ctx.ArgCode {
		msgStr = "```\n" + msgStr + "```\n"
	}

	if *ctx.ArgWhere == "slack" {
		println("slack")
		MsgSlack(&ctx, msgStr)
	} else if *ctx.ArgWhere == "telegram" {
		println("telegram")
		MsgTelegram(&ctx, msgStr)
	} else if *ctx.ArgWhere == "mailgun" {
		println("mailgun")
		MsgMailgun(&ctx, msgStr)
	} else {
		println("usage")
		flag.Usage()
	}
}

func getMsg() string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fn := strings.Trim(scanner.Text(), " \t")
		lines = append(lines, fn)
	}
	if scanner.Err() != nil {
		// handle error.
	}
	msgStr := strings.Join(lines, "\n")
	return msgStr
}

