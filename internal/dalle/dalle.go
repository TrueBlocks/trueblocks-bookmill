package dalle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type Response struct {
	Data []ImageData `json:"data"`
}

type ImageData struct {
	URL string `json:"url"`
}

func GenerateImage(apiKey, prompt, model, size, quality string) ([]byte, error) {
	reqBody := Request{
		Model:  model,
		Prompt: prompt,
		N:      1,
		Size:   size,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API call: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var dalleResp Response
	if err := json.Unmarshal(body, &dalleResp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	if len(dalleResp.Data) == 0 {
		return nil, fmt.Errorf("no images returned")
	}

	imageURL := dalleResp.Data[0].URL
	if imageURL == "" {
		return nil, fmt.Errorf("empty image URL in response")
	}

	imgClient := &http.Client{Timeout: 60 * time.Second}
	imgResp, err := imgClient.Get(imageURL)
	if err != nil {
		return nil, fmt.Errorf("downloading image: %w", err)
	}
	defer imgResp.Body.Close()

	if imgResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("image download error %d", imgResp.StatusCode)
	}

	imgData, err := io.ReadAll(imgResp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading image: %w", err)
	}

	return imgData, nil
}
