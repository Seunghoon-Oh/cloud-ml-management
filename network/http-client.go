package network

import (
	"time"

	circuit "github.com/rubyist/circuitbreaker"
)

func GetHttpClient() (*circuit.HTTPClient, *circuit.Breaker) {
	circuitBreaker := circuit.NewThresholdBreaker(5)
	client := circuit.NewHTTPClientWithBreaker(circuitBreaker, time.Millisecond*100, nil)
	return client, circuitBreaker
}
