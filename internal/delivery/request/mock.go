package request

import (
	"net/http"

	"github.com/elvenworks/sidekiq-conector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SidekiqMock struct {
	mock.Mock
}

func (m *SidekiqMock) Request(p domain.SidekiqParameters) ([]byte, *http.Response, error) {
	args := m.Called(p)
	return args.Get(0).([]byte), args.Get(1).(*http.Response), args.Error(2)
}
