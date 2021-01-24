package main

import (
	"fmt"
	"strings"

	"encoding/json"
	"net/http"
	"net/url"
)

func MsgSMS(ctx *Context, msg string) error {
	cfg := ctx.Config

	sid := cfg.Twilios[*ctx.ArgProfileName].Sid
	token := cfg.Twilios[*ctx.ArgProfileName].Token

	from := cfg.Twilios[*ctx.ArgProfileName].From
	if from == "" {
		from = *ctx.ArgFrom
	}
	to := cfg.Twilios[*ctx.ArgProfileName].To
	if to == "" {
		to = *ctx.ArgTo
	}

	fmt.Printf("To: %s From: %s\n", to, from)
	fmt.Printf("SID: %s TOKEN: %s MSG: %s\n", sid, token, msg)

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", msg)
	msgDataReader := *strings.NewReader(msgData.Encode())

	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", sid)

	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(sid, token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	client := &http.Client{}
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
	return nil
}
