package messenger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/wkoszek/msgr/internal/config"
)

// Twilio sends SMS messages via Twilio.
type Twilio struct {
	sid   string
	token string
	from  string
	to    string
}

// NewTwilio creates a new Twilio messenger.
func NewTwilio(cfg *config.Twilio, opts Options) *Twilio {
	from := cfg.From
	if from == "" {
		from = opts.From
	}
	to := cfg.To
	if to == "" {
		to = opts.To
	}

	return &Twilio{
		sid:   cfg.Sid,
		token: cfg.Token,
		from:  from,
		to:    to,
	}
}

// Name returns the messenger name.
func (t *Twilio) Name() string {
	return "twilio"
}

// Send sends an SMS via Twilio.
func (t *Twilio) Send(ctx context.Context, msg string) error {
	data := url.Values{}
	data.Set("To", t.to)
	data.Set("From", t.from)
	data.Set("Body", msg)

	apiURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", t.sid)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.SetBasicAuth(t.sid, t.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("twilio returned status %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}
