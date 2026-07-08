package config

import "log/slog"

type ConfigOption func(*Config) error

type Config struct {
	Port     string
	LogLevel slog.Level
}

func NewConfig(opts ...ConfigOption) (*Config, error) {
	c := &Config{}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithPort(port string) ConfigOption {
	return func(c *Config) error {
		c.Port = port
		return nil
	}
}

func WithLogLevel(level string) ConfigOption {
	return func(c *Config) error {
		var logLevel slog.Level
		if err := logLevel.UnmarshalText([]byte(level)); err != nil {
			return err
		}
		slog.SetLogLoggerLevel(logLevel)
		c.LogLevel = logLevel
		return nil
	}
}
