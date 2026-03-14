package pipeline

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type AnthropicClient struct {
	APIKey  string
	OnRetry func(attempt, maxAttempts int, err error, nextIn time.Duration)
}

type anthropicRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicResponse struct {
	Content []contentBlock `json:"content"`
	Usage   usage          `json:"usage"`
	Error   *apiError      `json:"error,omitempty"`
}

type contentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type apiError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type APIResult struct {
	Content      string
	InputTokens  int
	OutputTokens int
	Cost         float64
}

func (c *AnthropicClient) Call(ctx context.Context, model, prompt string, timeout time.Duration) (*APIResult, error) {
	const maxAttempts = 30
	const retryWait = 2 * time.Minute
	deadline := time.Now().Add(1 * time.Hour)

	for attempt := 0; ; attempt++ {
		result, err := c.callOnce(ctx, model, prompt, timeout)
		if err == nil {
			if attempt > 0 && c.OnRetry != nil {
				c.OnRetry(0, 0, nil, 0)
			}
			return result, nil
		}

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		if !isRetryableError(err) || attempt >= maxAttempts-1 || time.Now().After(deadline) {
			if c.OnRetry != nil {
				c.OnRetry(0, 0, nil, 0)
			}
			return nil, err
		}

		secs := int(retryWait.Seconds())
		for s := secs; s > 0; s-- {
			if c.OnRetry != nil {
				c.OnRetry(attempt+1, maxAttempts, err, time.Duration(s)*time.Second)
			}
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(1 * time.Second):
			}
		}
	}
}

func isRetryableError(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "credit balance") ||
		strings.Contains(msg, "rate_limit") ||
		strings.Contains(msg, "overloaded") ||
		strings.Contains(msg, "deadline exceeded")
}

func (c *AnthropicClient) callOnce(ctx context.Context, model, prompt string, timeout time.Duration) (*APIResult, error) {
	reqBody := anthropicRequest{
		Model:     model,
		MaxTokens: 8192,
		Messages: []message{
			{Role: "user", Content: prompt},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.anthropic.com/v1/messages", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("calling API: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var apiResp anthropicResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	if apiResp.Error != nil {
		return nil, fmt.Errorf("API error (%s): %s", apiResp.Error.Type, apiResp.Error.Message)
	}

	var text string
	for _, block := range apiResp.Content {
		if block.Type == "text" {
			text += block.Text
		}
	}

	cost := estimateCost(model, apiResp.Usage.InputTokens, apiResp.Usage.OutputTokens)

	return &APIResult{
		Content:      text,
		InputTokens:  apiResp.Usage.InputTokens,
		OutputTokens: apiResp.Usage.OutputTokens,
		Cost:         cost,
	}, nil
}

func estimateCost(model string, inputTokens, outputTokens int) float64 {
	var inputRate, outputRate float64
	switch {
	case strings.Contains(model, "opus"):
		inputRate = 15.0 / 1_000_000
		outputRate = 75.0 / 1_000_000
	default:
		inputRate = 3.0 / 1_000_000
		outputRate = 15.0 / 1_000_000
	}
	return float64(inputTokens)*inputRate + float64(outputTokens)*outputRate
}

func DryRunResult(stage Stage, title string) *APIResult {
	content := fmt.Sprintf("# DRY RUN: %s — %s\n\n"+
		"This is placeholder content generated in dry-run mode.\n\n"+
		"In live mode, this would contain the AI-generated %s output for \"%s\".\n\n"+
		"The prompt template and model configuration are ready.\n"+
		"Set `dry_run: false` in config.yaml and provide an API key to generate real content.\n",
		stage.String(), title, stage.String(), title)

	return &APIResult{
		Content:      content,
		InputTokens:  0,
		OutputTokens: 0,
		Cost:         0,
	}
}
