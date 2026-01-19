package messenger

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v3"

	"github.com/wkoszek/msgr/internal/config"
)

// Mailgun sends messages via Mailgun email service.
type Mailgun struct {
	domain  string
	key     string
	from    string
	to      string
	subject string
	html    bool
}

// NewMailgun creates a new Mailgun messenger.
func NewMailgun(cfg *config.Mailgun, opts Options) *Mailgun {
	return &Mailgun{
		domain:  cfg.Domain,
		key:     cfg.Key,
		from:    opts.From,
		to:      opts.To,
		subject: opts.Subject,
		html:    opts.HTML,
	}
}

// Name returns the messenger name.
func (m *Mailgun) Name() string {
	return "mailgun"
}

// Send sends an email via Mailgun.
func (m *Mailgun) Send(ctx context.Context, msg string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mg := mailgun.NewMailgun(m.domain, m.key)
	message := mg.NewMessage(m.from, m.subject, msg, m.to)

	if m.html {
		message.SetHtml(msg)
	}

	_, _, err := mg.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("sending email: %w", err)
	}

	return nil
}
