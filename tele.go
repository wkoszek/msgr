package main

import "fmt"
import "net/url"

// MsgTelegram sends a message via Telegram
func MsgTelegram(ctx *Context, msg string) error {
	cfg := ctx.Config
	token := cfg.Telegrams[*ctx.ArgProfileName].Token
	chatID := cfg.Telegrams[*ctx.ArgProfileName].ChatID

	println("telegram 2")

	fmt.Println(msg)

	msgEncoded := url.QueryEscape(msg)

	getReqStr :=
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&parse_mode=Markdown&text=%s",
			token, chatID, msgEncoded)

	err := httpGet(getReqStr)

	return err
}
