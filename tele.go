package main

import "fmt"

func MsgTelegram(cfg *Config, msg string) error {
	token := cfg.Telegrams[cfg.ProfileName].Token
	chatID := cfg.Telegrams[cfg.ProfileName].ChatID

	println("telegram 2")

	getReqStr := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&parse_mode=Markdown&text=%s", token, chatID, msg)

	err := httpGet(getReqStr)

	return err
}
