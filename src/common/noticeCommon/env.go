package noticeCommon

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type IGoogleChat interface {
	Send() error
}

func GoogleChatSend(url string, payload map[string]interface{}) error {

	payloadString, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadString))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil

}
