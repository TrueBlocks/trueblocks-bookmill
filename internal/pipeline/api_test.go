package pipeline

import (
	"fmt"
	"math"
	"testing"
)

func TestEstimateCost(t *testing.T) {
	t.Parallel()
	pricing := map[string]*ModelPricing{
		"sonnet":  {InputPer1M: 3.0, OutputPer1M: 15.0},
		"opus":    {InputPer1M: 15.0, OutputPer1M: 75.0},
		"default": {InputPer1M: 1.0, OutputPer1M: 5.0},
	}
	tests := []struct {
		name         string
		model        string
		input        int
		output       int
		wantApprox   float64
	}{
		{"sonnet match", "claude-3-sonnet-20240229", 1000, 500, 1000*3.0/1e6 + 500*15.0/1e6},
		{"opus match", "claude-3-opus-20240229", 1000, 500, 1000*15.0/1e6 + 500*75.0/1e6},
		{"default fallback", "unknown-model", 1000, 500, 1000*1.0/1e6 + 500*5.0/1e6},
		{"zero tokens", "claude-3-sonnet", 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := estimateCost(tt.model, tt.input, tt.output, pricing)
			if math.Abs(got-tt.wantApprox) > 1e-9 {
				t.Errorf("estimateCost() = %v, want %v", got, tt.wantApprox)
			}
		})
	}
}

func TestEstimateCostNilPricing(t *testing.T) {
	t.Parallel()
	got := estimateCost("anything", 1_000_000, 1_000_000, nil)
	want := 3.0 + 15.0
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("estimateCost(nil pricing) = %v, want %v", got, want)
	}
}

func TestIsRetryableError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg  string
		want bool
	}{
		{"insufficient credit balance", true},
		{"rate_limit exceeded", true},
		{"server overloaded", true},
		{"context deadline exceeded", true},
		{"invalid request", false},
		{"connection refused", false},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()
			if got := isRetryableError(fmt.Errorf("%s", tt.msg)); got != tt.want {
				t.Errorf("isRetryableError(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
