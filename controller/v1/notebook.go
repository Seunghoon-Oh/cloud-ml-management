package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Seunghoon-Oh/cloud-ml-manager/service"
)

// func GetNotebooks() {
// 	// GET 호출
// 	resp, err := http.Get("http://csharp.news")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer resp.Body.Close()

// 	// 결과 출력
// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", string(data))
// }

func GetNotebooks(c *gin.Context) {
	data := service.GetNotebooks()
	println(data)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
