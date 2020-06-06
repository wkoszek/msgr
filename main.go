package main

import "fmt"
import "flag"
import "os"
import "os/user"
import "bufio"
import "strings"
import "path"


func main() {
	var	ctx	Context

	ctx.ArgWhere = flag.String("where", "slack", "service to post [slack]")
	ctx.ArgProfileName = flag.String("profile", "default", "profile name [default]")
	ctx.ArgCode = flag.Bool("code", false, "paste code block [disabled]")
	ctx.ArgHTML = flag.Bool("html", false, "use HTML for emails")
	ctx.ArgFrom = flag.String("from", "default", "From field (email)")
	ctx.ArgTo = flag.String("to", "default", "To field (email)")
	ctx.ArgSubject = flag.String("subject", "default", "Subject field (email)")
	ctx.ArgDry = flag.Bool("dry", false, "Dry mode (won't send anything)")

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

	configPath := getConfigPath()
	cfg := NewConfig(configPath, *ctx.ArgProfileName)
	ctx.Config = cfg

	// @todo think of doing it so that listening for stdin could be
	// continuous. if there's 100s delay and something gets spitted
	// out, maybe we could send it to slack too
	msgStr := getMsg()

	// @todo adjust for HTML
	if *ctx.ArgCode {
		msgStr = "```\n" + msgStr + "```\n"
	}

	if *ctx.ArgDry {
		fmt.Printf("# Dry mode, this would have been sent, but will be muffled\n")
		fmt.Printf("%s\n", msgStr)
		os.Exit(0)
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
	reader := bufio.NewReader(os.Stdin)

	fullText := ""
	for {
		txt, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fullText += (strings.Trim(txt, "\n") + "\n")
	}

	return fullText
}

func getConfigPath() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	// @todo lookup what the modern standards are for config files

	configPathAlt1 := path.Join(u.HomeDir, ".config", ".msgr.conf")
	configPathAlt2 := path.Join(u.HomeDir, ".msgr.conf")

	fmt.Printf("files:", configPathAlt1, configPathAlt2)

	configPath := configPathAlt1
	_, err = os.Stat(configPathAlt1)
	if os.IsNotExist(err) {
		_, err = os.Stat(configPathAlt2)
		configPath = configPathAlt2
	}
	if os.IsNotExist(err) {
		panic("can't find config file")
	}
	return configPath
}
