package service

import (
	"io"
	"net/http"
)

func GetStudios() string {
	resp, err := http.Get("http://cloud-ml-studio-manager.cloud-ml-studio:8082/studios")
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
