package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var fooCb *circuit.Breaker
var fooClienct *circuit.HTTPClient

func SetupFooCircuitBreaker() {
	fooClienct, fooCb = network.GetHttpClient()
}

func GetFoos() []string {
	if fooCb.Ready() {
		resp, err := fooClienct.Get("http://cloud-ml-foo-manager.cloud-ml-foo:8082/foos")
		if err != nil {
			fmt.Println(err)
			fooCb.Fail()
			return nil
		}
		fooCb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		return rsData.Data
	}
	return nil
}
