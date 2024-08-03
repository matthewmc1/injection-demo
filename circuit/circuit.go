package circuit

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type BreakerConfig struct {
	MaxRetries   int    `yaml:"retries"`
	Backoff      int    `yaml:"backoff"`
	MaxThreshold int    `yaml:"threshold"`
	URL          string `yaml:"url"`
}

type BreakerInterface interface {
	DefaultHandler(handler func(c context.Context, w http.ResponseWriter, r *http.Request))
	HandlerWithDeadline()
}

func (bc *BreakerConfig) Breaker(hand func(c context.Context, url string)) func(c context.Context, url string) (int, error) {
	return func(c context.Context, url string) (int, error) {
		var i int
		r, _ := http.Get(url)
		if r.StatusCode != 200 {
			for i = 1; i <= bc.MaxRetries; i++ {
				r, _ := http.Get(url)
				if r.StatusCode == 200 {
					return r.StatusCode, nil
				}
				time.Sleep(time.Duration(bc.Backoff))
			}
			return r.StatusCode, errors.New("reached max backoff")
		}
		return r.StatusCode, nil
	}
}

func Request(c context.Context, url string) {
	if url == "" {
		return
	}
}
