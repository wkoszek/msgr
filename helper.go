package main;

import "net/http"
import "log"

func httpGet(urlStr string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return err
}
