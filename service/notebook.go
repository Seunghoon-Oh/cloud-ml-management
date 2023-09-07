package service

import (
	"fmt"
	"io"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var notebookCb *circuit.Breaker
var notebookClient *circuit.HTTPClient

func SetupNotebookCircuitBreaker() {
	notebookClient, notebookCb = network.GetHttpClient()
}

func GetNotebooks() string {
	var result string
	if notebookCb.Ready() {
		resp, err := notebookClient.Get("http://cloud-ml-notebook-manager.cloud-ml-notebook:8082/notebooks")
		if err != nil {
			fmt.Println(err)
			notebookCb.Fail()
			return result
		}
		notebookCb.Success()
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
