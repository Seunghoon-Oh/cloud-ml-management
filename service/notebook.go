package service

import (
	"io"
	"net/http"
)

func GetNotebooks() string {
	resp, err := http.Get("http://cloud-ml-notebook-manager.cloud-ml-notebook:8082/notebooks")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(data)
}
