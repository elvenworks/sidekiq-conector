package domain

type sidekiqInfos struct {
	Busy      int `json:"busy"`
	Enqueued  int `json:"enqueued"`
	Scheduled int `json:"scheduled"`
}

type Sidekiq struct {
	Sidekiq *sidekiqInfos `json:"sidekiq"`
	Redis   *sidekiqRedis `json:"redis"`
}
