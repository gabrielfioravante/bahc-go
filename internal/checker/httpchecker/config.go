package httpchecker

import "time"

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST            = "POST"
)

type HttpCheckerConfig struct {
	ID            string        `yaml:"id"`
	Url           string        `yaml:"url"`
	Method        HttpMethod    `yaml:"host"`
	Authorization string        `yaml:"authorization"`
	Timeout       time.Duration `yaml:"timeout"`
	Interval      time.Duration `yaml:"interval"`
}
