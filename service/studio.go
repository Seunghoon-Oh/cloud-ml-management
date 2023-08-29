package service

import (
	"fmt"
	"io"
	"net/http"
)

func GetStudios() string {
	resp, err := http.Get("http://cloud-ml-studio-manager.cloud-ml-studio:8082/")
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
