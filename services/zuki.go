package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ZukiInterface struct {
	Auth    string
	BaseURL string
}

var zuki *ZukiInterface

func InitZuki(auth string) {
	zuki = &ZukiInterface{
		Auth:    auth,
		BaseURL: "https://zukijourney.xyzbot.net",
	}
}

func ChatCall(messages []map[string]string, model string) ([]byte, error) {
	body := map[string]interface{}{
		"model":    model,
		"messages": messages,
	}
	return zuki.makeRequest("/v1/chat/completions", body)
}

func ImageCall(model, prompt, size, quality string, n int) ([]byte, error) {
	body := map[string]interface{}{
		"prompt":  prompt,
		"model":   model,
		"size":    size,
		"quality": quality,
		"n":       n,
	}
	return zuki.makeRequest("/v1/images/generations", body)
}

func (z *ZukiInterface) makeRequest(endpoint string, data map[string]interface{}) ([]byte, error) {
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", z.BaseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+z.Auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}
