package sidekiq

import (
	"github.com/elvenworks/sidekiq-conector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SidekiqMock struct {
	mock.Mock
}

func (m *SidekiqMock) SidekiqReq() (sidekiqReponse *domain.Sidekiq, err error) {
	args := m.Called()
	return args.Get(0).(*domain.Sidekiq), args.Error(1)
}
