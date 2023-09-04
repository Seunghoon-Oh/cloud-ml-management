package service

import (
	"fmt"
	"io"

	"github.com/Seunghoon-Oh/cloud-ml-manager/network"
)

// func GetStudios() string {
// 	client := network.GetHttpClient()
// 	resp, err := client.Get("http://cloud-ml-studio-manager.cloud-ml-studio:8082/studios")
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

func GetStudios() string {
	var result string
	client, cb := network.GetHttpClient()
	if cb.Ready() {
		resp, err := client.Get("http://cloud-ml-studio-manager.cloud-ml-studio:8082/studios")
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
