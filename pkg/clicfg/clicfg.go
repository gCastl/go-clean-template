package clicfg

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/pflag"
)

type Field struct {
	EnvVar     string
	DefaultVal string
	Value      string
	Type       reflect.Kind
}

type Config struct {
	fields map[string]*Field
}

type Option func(*Config)

func (c *Config) GetString(varName string) string {
	field := c.fields[varName]
	if field == nil {
		panic(fmt.Sprintf("config field %q not found", varName))
	}
	return field.Value
}

func (c *Config) GetBool(varName string) bool {
	field := c.fields[varName]
	if field == nil {
		panic(fmt.Sprintf("config field %q not found", varName))
	}
	val, err := strconv.ParseBool(field.Value)
	if err != nil {
		panic(fmt.Sprintf("config field %q: invalid bool value %q: %v", varName, field.Value, err))
	}
	return val
}

func (c *Config) GetInt(varName string) int {
	field := c.fields[varName]
	if field == nil {
		panic(fmt.Sprintf("config field %q not found", varName))
	}
	val, err := strconv.Atoi(field.Value)
	if err != nil {
		panic(fmt.Sprintf("config field %q: invalid int value %q: %v", varName, field.Value, err))
	}
	return val
}

func (c *Config) GetFloat(varName string) float64 {
	field := c.fields[varName]
	if field == nil {
		panic(fmt.Sprintf("config field %q not found", varName))
	}
	val, err := strconv.ParseFloat(field.Value, 64)
	if err != nil {
		panic(fmt.Sprintf("config field %q: invalid float value %q: %v", varName, field.Value, err))
	}
	return val
}

func New(opts ...Option) *Config {
	c := &Config{
		fields: make(map[string]*Field),
	}

	for _, opt := range opts {
		opt(c)
	}

	flagSet := pflag.NewFlagSet("app", pflag.ContinueOnError)

	// Register all flags with pflag
	for fieldName, field := range c.fields {
		switch field.Type {
		case reflect.String:
			flagSet.String(fieldName, field.DefaultVal, "")
		case reflect.Bool:
			defaultBool, _ := strconv.ParseBool(field.DefaultVal)
			flagSet.Bool(fieldName, defaultBool, "")
		case reflect.Int64:
			defaultInt, _ := strconv.ParseInt(field.DefaultVal, 10, 64)
			flagSet.Int64(fieldName, defaultInt, "")
		case reflect.Float64:
			defaultFloat, _ := strconv.ParseFloat(field.DefaultVal, 64)
			flagSet.Float64(fieldName, defaultFloat, "")
		}
	}

	// Parse command-line flags
	flagSet.Parse(os.Args[1:])

	// Resolve values: flag > env var > default
	for fieldName, field := range c.fields {
		fl := flagSet.Lookup(fieldName)
		field.Value = field.DefaultVal

		if fl == nil {
			continue
		}

		// Check if flag was explicitly provided
		if fl.Changed {
			field.Value = fl.Value.String()
			continue
		}

		// Check environment variable
		if env, ok := os.LookupEnv(field.EnvVar); ok {
			field.Value = env
		}
	}

	return c
}

func String(flagName, envVar, defaultValue string) Option {
	return func(c *Config) {
		c.fields[flagName] = &Field{
			EnvVar:     envVar,
			DefaultVal: defaultValue,
			Type:       reflect.String,
		}
	}
}

func Bool(flagName, envVar string, defaultValue bool) Option {
	return func(c *Config) {
		c.fields[flagName] = &Field{
			EnvVar:     envVar,
			DefaultVal: strconv.FormatBool(defaultValue),
			Type:       reflect.Bool,
		}
	}
}

func Int(flagName, envVar string, defaultValue int64) Option {
	return func(c *Config) {
		c.fields[flagName] = &Field{
			EnvVar:     envVar,
			DefaultVal: strconv.FormatInt(defaultValue, 10),
			Type:       reflect.Int64,
		}
	}
}

func Float(flagName, envVar string, defaultValue float64) Option {
	return func(c *Config) {
		c.fields[flagName] = &Field{
			EnvVar:     envVar,
			DefaultVal: strconv.FormatFloat(defaultValue, 'f', -1, 64),
			Type:       reflect.Float64,
		}
	}
}
