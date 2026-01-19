package messenger

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/wkoszek/msgr/internal/config"
)

// Telegram sends messages via Telegram Bot API.
type Telegram struct {
	token  string
	chatID string
}

// NewTelegram creates a new Telegram messenger.
func NewTelegram(cfg *config.Telegram) *Telegram {
	return &Telegram{
		token:  cfg.Token,
		chatID: cfg.ChatID,
	}
}

// Name returns the messenger name.
func (t *Telegram) Name() string {
	return "telegram"
}

// Send sends a message via Telegram.
func (t *Telegram) Send(ctx context.Context, msg string) error {
	encodedMsg := url.QueryEscape(msg)
	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage?chat_id=%s&parse_mode=Markdown&text=%s",
		t.token, t.chatID, encodedMsg,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram returned status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
