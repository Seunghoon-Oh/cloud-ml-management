package service

import (
	"fmt"
	"io"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var pipelineCb *circuit.Breaker
var pipelineClienct *circuit.HTTPClient

func SetupPipelineCircuitBreaker() {
	pipelineClienct, pipelineCb = network.GetHttpClient()
}

func GetPipelines() string {
	var result string

	if pipelineCb.Ready() {
		resp, err := pipelineClienct.Get("http://cloud-ml-pipeline-manager.cloud-ml-pipeline:8082/pipelines")
		if err != nil {
			fmt.Println(err)
			pipelineCb.Fail()
			return result
		}
		pipelineCb.Success()
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
