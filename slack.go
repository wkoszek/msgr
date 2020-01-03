package main

import "bytes"
import "encoding/json"
import "errors"
import "net/http"
import "time"
import "log"

type SlackRequestBody struct {
	Text string `json:"text"`
}

func MsgSlack(cfg *Config, msg string) error {
	webhookUrl := "https://hooks.slack.com/services/" + cfg.Slacks[cfg.ProfileName].Key

	log.Printf("webhook: %s", webhookUrl)

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
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
