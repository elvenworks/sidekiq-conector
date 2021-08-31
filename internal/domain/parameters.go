package domain

import (
	"time"
)

type SidekiqParameters struct {
	Url               string
	SkipSSLValidation bool
	Headers           map[string]interface{}
	Timeout           time.Duration
}
