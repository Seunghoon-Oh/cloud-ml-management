package service

import (
	"fmt"
	"io"
	"net/http"
)

func GetNotebooks() string {

	// 	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	//필요시 헤더 추가 가능
	// 	req.Header.Add("User-Agent", "Crawler")

	// 	// Client객체에서 Request 실행
	// 	client := &http.Client{}
	// 	resp, err := client.Do(req)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer resp.Body.Close()

	// 	// 결과 출력
	// 	bytes, _ := io.ReadAll(resp.Body)
	// 	str := string(bytes) //바이트를 문자열로
	// 	fmt.Println(str)
	// 	c.String(200, "ok")
	// })

	// GET 호출
	resp, err := http.Get("http://cloud-ml-notebook-manager.cloud-ml-notebook:8082/")
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
