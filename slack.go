// Copyright (c) 2020 by Wojciech Adam Koszek <wojciech@koszek.com>
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type slackRequestBody struct {
	Text string `json:"text"`
}

// MsgSlack sends a message on Slack
func MsgSlack(ctx *Context, msg string) error {
	cfg := ctx.Config
	webhookURL := "https://hooks.slack.com/services/" + cfg.Slacks[*ctx.ArgProfileName].Key

	log.Printf("webhook: %s", webhookURL)

	slackBody, _ := json.Marshal(slackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
