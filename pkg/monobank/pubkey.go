package monobank

import (
	"encoding/json"
	"net/http"
)

const (
	pubKeyUrl = "https://api.monobank.ua/personal/public"
)

type PubKeyResponse struct {
	Key string `json:"key"`
}

func (c *Client) getPubKey() (string, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", pubKeyUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Token", c.xToken)
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var response PubKeyResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}
	return response.Key, nil
}
