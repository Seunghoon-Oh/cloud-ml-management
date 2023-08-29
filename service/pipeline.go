package service

import (
	"fmt"
	"io"
	"net/http"
)

func GetPipelines() string {
	resp, err := http.Get("http://cloud-ml-pipeline-manager.cloud-ml-pipeline:8082/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
	return string(data)
}
