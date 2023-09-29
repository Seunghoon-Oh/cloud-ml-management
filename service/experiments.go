package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var expCb *circuit.Breaker
var expClient *circuit.HTTPClient

func SetupExpCircuitBreaker() {
	expClient, expCb = network.GetHttpClient()
}

func GetExps() []string {
	if expCb.Ready() {
		resp, err := expClient.Get("http://cloud-ml-experiments-manager.cloud-ml-experiments:8082/exps")
		if err != nil {
			fmt.Println(err)
			expCb.Fail()
			return nil
		}
		expCb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		return rsData.Data
	}
	return nil
}
