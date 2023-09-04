package service

import (
	"fmt"
	"io"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
)

func GetNotebooks() string {
	var result string
	client, cb := network.GetHttpClient()
	if cb.Ready() {
		resp, err := client.Get("http://cloud-ml-notebook-manager.cloud-ml-notebook:8082/notebooks")
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
