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

func GetAllMLServervices(c *gin.Context) {

	nobooks := service.GetNotebooks()
	println(nobooks)
	studios := service.GetStudios()
	println(studios)
	pipelines := service.GetPipelines()
	println(pipelines)

	data := make([]string, 3)
	data = append(data, nobooks, studios, pipelines)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
