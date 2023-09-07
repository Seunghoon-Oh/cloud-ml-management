package service

import (
	"fmt"
	"io"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var studioCb *circuit.Breaker
var studioClienct *circuit.HTTPClient

func SetupStudioCircuitBreaker() {
	studioClienct, studioCb = network.GetHttpClient()
}

func GetStudios() string {
	var result string
	if studioCb.Ready() {
		resp, err := studioClienct.Get("http://cloud-ml-studio-manager.cloud-ml-studio:8082/studios")
		if err != nil {
			fmt.Println(err)
			studioCb.Fail()
			return result
		}
		studioCb.Success()
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		result = string(data)
		return result
	}
	return result
}
