package main

import "github.com/mailgun/mailgun-go/v3"
import "fmt"
import "context"
import "time"

// MsgEmailMailgun
func MsgMailgun(ctx *Context, msg string) error {
	cfg := ctx.Config

	domain := cfg.Mailguns[*ctx.ArgProfileName].Domain
	key := cfg.Mailguns[*ctx.ArgProfileName].Key

	mailCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mg := mailgun.NewMailgun(domain, key)
	message := mg.NewMessage(*ctx.ArgFrom, *ctx.ArgSubject, msg, *ctx.ArgTo) // this is text body
	if *ctx.ArgHTML {
		message.SetHtml(msg) // this is HTML body
	}
	resp, id, err := mg.Send(mailCtx, message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
