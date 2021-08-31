package request

import (
	"net/http"

	"github.com/elvenworks/sidekiq-conector/internal/domain"
)

type ISidekiq interface {
	Request(p domain.SidekiqParameters) ([]byte, *http.Response, error)
}
