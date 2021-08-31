package sidekiq

import "github.com/elvenworks/sidekiq-conector/internal/domain"

type ISidekiq interface {
	SidekiqReq() (sidekiqReponse *domain.Sidekiq, err error)
}
