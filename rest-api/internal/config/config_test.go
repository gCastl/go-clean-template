package config

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestNewConfig_NoOptions(t *testing.T) {
	c, err := NewConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil config")
	}
}

func TestNewConfig_WithLogLevel(t *testing.T) {
	tests := []struct {
		level    string
		expected slog.Level
		wantErr  bool
	}{
		{"DEBUG", slog.LevelDebug, false},
		{"INFO", slog.LevelInfo, false},
		{"WARN", slog.LevelWarn, false},
		{"ERROR", slog.LevelError, false},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			c, err := NewConfig(WithLogLevel(tt.level))
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if c.LogLevel != tt.expected {
				t.Errorf("expected log level %v, got %v", tt.expected, c.LogLevel)
			}
		})
	}
}

func TestNewConfig_OptionError(t *testing.T) {
	failOpt := func(*Config) error {
		return errTest
	}
	_, err := NewConfig(failOpt)
	if err == nil {
		t.Fatal("expected error from failing option, got nil")
	}
}

var errTest = fmt.Errorf("test error")
