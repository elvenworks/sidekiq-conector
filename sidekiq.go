package sidekiq

import (
	"encoding/json"
	"time"

	"github.com/elvenworks/sidekiq-conector/internal/delivery/request"
	"github.com/elvenworks/sidekiq-conector/internal/domain"
	"github.com/sirupsen/logrus"
)

type SidekiqConfig struct {
	Url               string
	SkipSSLValidation bool
	Headers           map[string]interface{}
	Timeout           time.Duration
}

type Sidekiq struct {
	sidekiq           request.ISidekiq
	sidekiqParameters domain.SidekiqParameters
}

func InitSidekiq(c SidekiqConfig) *Sidekiq {
	parameters := domain.SidekiqParameters{
		Url:               c.Url,
		SkipSSLValidation: c.SkipSSLValidation,
		Headers:           c.Headers,
		Timeout:           c.Timeout,
	}
	return &Sidekiq{
		sidekiqParameters: parameters,
		sidekiq:           request.NewSidekiq(),
	}
}

func (p Sidekiq) SidekiqReq() (sidekiqReponse *domain.Sidekiq, err error) {

	contents, _, err := p.sidekiq.Request(p.sidekiqParameters)

	if err != nil {
		return nil, err
	}

	var info domain.Sidekiq

	if err := json.Unmarshal(contents, &info); err != nil {
		logrus.Errorf("Failed to return sidekiq infos, err: %s\n", err)
		return nil, err
	}

	return &info, nil
}
