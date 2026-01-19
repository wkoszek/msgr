package messenger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/wkoszek/msgr/internal/config"
)

// Slack sends messages via Slack webhooks.
type Slack struct {
	webhookKey string
}

// NewSlack creates a new Slack messenger.
func NewSlack(cfg *config.Slack) *Slack {
	return &Slack{webhookKey: cfg.Key}
}

// Name returns the messenger name.
func (s *Slack) Name() string {
	return "slack"
}

// Send sends a message to Slack.
func (s *Slack) Send(ctx context.Context, msg string) error {
	webhookURL := "https://hooks.slack.com/services/" + s.webhookKey

	body := struct {
		Text string `json:"text"`
	}{Text: msg}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshaling message: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, webhookURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}

	if string(respBody) != "ok" {
		return fmt.Errorf("slack returned: %s", string(respBody))
	}

	return nil
}
