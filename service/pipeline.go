package service

import (
	"fmt"
	"io"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
)

// func GetPipelines() string {
// 	client := network.GetHttpClient()
// 	resp, err := client.Get("http://cloud-ml-pipeline-manager.cloud-ml-pipeline:8082/pipelines")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer resp.Body.Close()

// 	// 결과 출력
// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(data)
// }

func GetPipelines() string {
	var result string
	client, cb := network.GetHttpClient()
	if cb.Ready() {
		resp, err := client.Get("http://cloud-ml-pipeline-manager.cloud-ml-pipeline:8082/pipelines")
		fmt.Print("resp: ")
		fmt.Println(resp)
		if err != nil {
			fmt.Print("err: ")
			fmt.Println(err)
			cb.Fail()
			return result
		}
		cb.Success()
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
