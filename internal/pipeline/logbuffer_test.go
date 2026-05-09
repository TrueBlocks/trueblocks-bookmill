package pipeline

import (
	"bytes"
	"io"
	"strings"
	"sync"
	"testing"
)

func TestLogBufferLevelClassification(t *testing.T) {
	var sink bytes.Buffer
	lb := NewLogBuffer(&sink, 100)

	cases := []struct {
		msg   string
		level string
	}{
		{"plain info\n", "info"},
		{"this is an ERROR\n", "error"},
		{"VERBOSE detail\n", "verbose"},
		{"both ERROR and VERBOSE — ERROR wins\n", "error"},
	}
	for _, c := range cases {
		if _, err := lb.Write([]byte(c.msg)); err != nil {
			t.Fatalf("write: %v", err)
		}
	}

	entries := lb.Entries(true)
	if len(entries) != len(cases) {
		t.Fatalf("expected %d entries, got %d", len(cases), len(entries))
	}
	for i, c := range cases {
		if entries[i].Level != c.level {
			t.Errorf("entry %d: expected level=%s, got %s (msg=%q)", i, c.level, entries[i].Level, entries[i].Message)
		}
		if strings.HasSuffix(entries[i].Message, "\n") {
			t.Errorf("entry %d: trailing newline not stripped: %q", i, entries[i].Message)
		}
	}
}

func TestLogBufferEntriesFiltersVerbose(t *testing.T) {
	var sink bytes.Buffer
	lb := NewLogBuffer(&sink, 100)

	_, _ = lb.Write([]byte("info one\n"))
	_, _ = lb.Write([]byte("VERBOSE noisy\n"))
	_, _ = lb.Write([]byte("info two\n"))

	all := lb.Entries(true)
	if len(all) != 3 {
		t.Fatalf("verbose=true: expected 3, got %d", len(all))
	}
	visible := lb.Entries(false)
	if len(visible) != 2 {
		t.Fatalf("verbose=false: expected 2, got %d", len(visible))
	}
	for _, e := range visible {
		if e.Level == "verbose" {
			t.Errorf("verbose entry leaked when filtered: %+v", e)
		}
	}
}

func TestLogBufferRingTruncation(t *testing.T) {
	var sink bytes.Buffer
	lb := NewLogBuffer(&sink, 3)

	for i, m := range []string{"a\n", "b\n", "c\n", "d\n", "e\n"} {
		if _, err := lb.Write([]byte(m)); err != nil {
			t.Fatalf("write %d: %v", i, err)
		}
	}
	entries := lb.Entries(true)
	if len(entries) != 3 {
		t.Fatalf("expected 3 entries after truncation, got %d", len(entries))
	}
	if entries[0].Message != "c" || entries[1].Message != "d" || entries[2].Message != "e" {
		t.Fatalf("expected oldest entries dropped, got %+v", entries)
	}
}

func TestLogBufferVersionIncrements(t *testing.T) {
	var sink bytes.Buffer
	lb := NewLogBuffer(&sink, 100)

	if lb.Version() != 0 {
		t.Fatalf("expected initial version 0, got %d", lb.Version())
	}
	for i := 0; i < 5; i++ {
		_, _ = lb.Write([]byte("x\n"))
	}
	if got := lb.Version(); got != 5 {
		t.Fatalf("expected version 5, got %d", got)
	}
}

func TestLogBufferWritesPassThrough(t *testing.T) {
	var sink bytes.Buffer
	lb := NewLogBuffer(&sink, 100)

	for _, m := range []string{"hello\n", "world\n"} {
		if _, err := lb.Write([]byte(m)); err != nil {
			t.Fatalf("write: %v", err)
		}
	}
	got := sink.String()
	if got != "hello\nworld\n" {
		t.Fatalf("unexpected sink contents: %q", got)
	}
}

func TestLogBufferConcurrentSafe(t *testing.T) {
	// Use io.Discard so the race detector exercises LogBuffer's own
	// synchronization, not bytes.Buffer (which is not goroutine-safe).
	lb := NewLogBuffer(io.Discard, 1000)

	var wg sync.WaitGroup
	const writers = 10
	const perWriter = 100
	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < perWriter; j++ {
				_, _ = lb.Write([]byte("x\n"))
			}
		}()
	}
	wg.Wait()

	if got := lb.Version(); got != writers*perWriter {
		t.Fatalf("expected version %d, got %d", writers*perWriter, got)
	}
	if entries := lb.Entries(true); len(entries) != writers*perWriter {
		t.Fatalf("expected %d entries, got %d", writers*perWriter, len(entries))
	}
}
