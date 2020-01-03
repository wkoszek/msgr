package main

import "fmt"
import "flag"
import "os"
import "bufio"
import "strings"

var ArgWhere = flag.String("where", "slack", "service to post [slack]")
var ArgCode = flag.Bool("code", false, "paste code block [disabled]")
var ArgProfile = flag.String("profile", "default", "profile name [default]")

func main() {
	flag.Parse()

	if *ArgWhere != "slack" && *ArgWhere != "telegram" {
		fmt.Println("%s unsupported", *ArgWhere)
		os.Exit(1)
	}

	println("before config")
	cfg := NewConfig(".msgr.conf", *ArgProfile)

	println("before getMsg")
	msgStr := getMsg()

	if *ArgCode {
		msgStr = "```" + msgStr + "```"
	}

	if *ArgWhere == "slack" {
		println("slack")
		MsgSlack(cfg, msgStr)
	} else if *ArgWhere == "telegram" {
		println("telegram")
		MsgTelegram(cfg, msgStr)
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

