package messenger

import (
	"context"
	"fmt"

	"github.com/wkoszek/msgr/internal/config"
)

// Messenger is the interface for sending messages.
type Messenger interface {
	Send(ctx context.Context, msg string) error
	Name() string
}

// Options contains optional settings for message delivery.
type Options struct {
	From    string
	To      string
	Subject string
	HTML    bool
}

// New creates a new Messenger based on the service name.
func New(service, profile string, cfg *config.Config, opts Options) (Messenger, error) {
	switch service {
	case "slack":
		slackCfg, ok := cfg.Slacks[profile]
		if !ok {
			return nil, fmt.Errorf("slack profile %q not found", profile)
		}
		return NewSlack(slackCfg), nil

	case "telegram":
		teleCfg, ok := cfg.Telegrams[profile]
		if !ok {
			return nil, fmt.Errorf("telegram profile %q not found", profile)
		}
		return NewTelegram(teleCfg), nil

	case "mailgun":
		mgCfg, ok := cfg.Mailguns[profile]
		if !ok {
			return nil, fmt.Errorf("mailgun profile %q not found", profile)
		}
		return NewMailgun(mgCfg, opts), nil

	case "twilio":
		twCfg, ok := cfg.Twilios[profile]
		if !ok {
			return nil, fmt.Errorf("twilio profile %q not found", profile)
		}
		return NewTwilio(twCfg, opts), nil

	default:
		return nil, fmt.Errorf("unknown service: %s", service)
	}
}

// SupportedServices returns the list of supported messaging services.
func SupportedServices() []string {
	return []string{"slack", "telegram", "mailgun", "twilio"}
}
