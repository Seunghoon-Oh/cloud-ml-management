package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var notebookCb *circuit.Breaker
var notebookClient *circuit.HTTPClient

func SetupNotebookCircuitBreaker() {
	notebookClient, notebookCb = network.GetHttpClient()
}

func GetNotebooks() []string {
	if notebookCb.Ready() {
		resp, err := notebookClient.Get("http://cloud-ml-notebook-manager.cloud-ml-notebook:8082/notebooks")
		if err != nil {
			fmt.Println(err)
			notebookCb.Fail()
			return nil
		}
		notebookCb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		return rsData.Data
	}
	return nil
}
