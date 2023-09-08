package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var pipelineCb *circuit.Breaker
var pipelineClienct *circuit.HTTPClient

func SetupPipelineCircuitBreaker() {
	pipelineClienct, pipelineCb = network.GetHttpClient()
}

func GetPipelines() []string {
	if pipelineCb.Ready() {
		resp, err := pipelineClienct.Get("http://cloud-ml-pipeline-manager.cloud-ml-pipeline:8082/pipelines")
		if err != nil {
			fmt.Println(err)
			pipelineCb.Fail()
			return nil
		}
		pipelineCb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		return rsData.Data
	}
	return nil
}
