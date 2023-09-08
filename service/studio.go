package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var studioCb *circuit.Breaker
var studioClienct *circuit.HTTPClient

func SetupStudioCircuitBreaker() {
	studioClienct, studioCb = network.GetHttpClient()
}

func GetStudios() []string {
	if studioCb.Ready() {
		resp, err := studioClienct.Get("http://cloud-ml-studio-manager.cloud-ml-studio:8082/studios")
		if err != nil {
			fmt.Println(err)
			studioCb.Fail()
			return nil
		}
		studioCb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		return rsData.Data
	}
	return nil
}
