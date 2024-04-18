package wclinet

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendMessage(webhookURL, message string) error {
	payload, err := json.Marshal(map[string]string{"content": message})
	if err != nil {
		return err
	}

	_, err = http.Post(webhookURL, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}

	return nil
}
