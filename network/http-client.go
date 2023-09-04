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

// func main() {
// 	circuitBreaker := circuit.NewThresholdBreaker(10)
// 	BreakerEvent := circuitBreaker.Subscribe()
// 	go func() {
// 		for {
// 			<-BreakerEvent
// 		}
// 	}()
// 	for {
// 		if circuitBreaker.Ready() {
// 			httpClient := circuit.NewHTTPClient(time.Second*10, 10, nil)
// 			resp, err := httpClient.Get("http://example1.comm")
// 			fmt.Println(resp)
// 			if err != nil {
// 				fmt.Println(err)
// 				circuitBreaker.Fail()
// 				continue
// 			}
// 			circuitBreaker.Success()
// 		}
// 	}
// }
