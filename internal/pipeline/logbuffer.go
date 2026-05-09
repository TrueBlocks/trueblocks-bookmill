package pipeline

import (
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type LogEntry struct {
	Time    string `json:"time"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

type LogBuffer struct {
	mu      sync.Mutex
	entries []LogEntry
	maxSize int
	out     io.Writer
	version atomic.Int64
}

func NewLogBuffer(out io.Writer, maxSize int) *LogBuffer {
	return &LogBuffer{out: out, maxSize: maxSize}
}

func (lb *LogBuffer) Write(p []byte) (n int, err error) {
	n, err = lb.out.Write(p)
	msg := string(p)
	if len(msg) > 0 && msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	level := "info"
	if strings.Contains(msg, "ERROR") {
		level = "error"
	} else if strings.Contains(msg, "VERBOSE") {
		level = "verbose"
	}
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.entries = append(lb.entries, LogEntry{
		Time:    time.Now().Format("15:04:05"),
		Message: msg,
		Level:   level,
	})
	if len(lb.entries) > lb.maxSize {
		lb.entries = lb.entries[len(lb.entries)-lb.maxSize:]
	}
	lb.version.Add(1)
	return
}

func (lb *LogBuffer) Entries(verbose bool) []LogEntry {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if verbose {
		out := make([]LogEntry, len(lb.entries))
		copy(out, lb.entries)
		return out
	}
	var out []LogEntry
	for _, e := range lb.entries {
		if e.Level != "verbose" {
			out = append(out, e)
		}
	}
	return out
}

func (lb *LogBuffer) Version() int64 {
	return lb.version.Load()
}
