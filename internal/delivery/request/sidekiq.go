package request

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"

	"github.com/elvenworks/sidekiq-conector/internal/domain"
)

type Sidekiq struct {
}

func NewSidekiq() *Sidekiq {
	return &Sidekiq{}
}

func (s *Sidekiq) Request(p domain.SidekiqParameters) ([]byte, *http.Response, error) {
	var err error
	var req *http.Request
	if req, err = http.NewRequest("GET", p.Url, nil); err != nil {
		return nil, nil, err
	}
	for key, value := range p.Headers {
		req.Header.Set(key, value.(string))
	}

	var resp *http.Response

	transport := &http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: p.Timeout,
		TLSHandshakeTimeout:   p.Timeout,
		Proxy:                 http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: p.SkipSSLValidation,
		},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   p.Timeout,
	}

	if resp, err = client.Do(req); err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	return contents, resp, err
}
