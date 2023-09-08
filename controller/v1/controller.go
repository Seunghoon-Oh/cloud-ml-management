package v1

import (
	"net/http"

	"github.com/Seunghoon-Oh/cloud-ml-manager/service"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Data string
}

func GetAllMLServervices(c *gin.Context) {

	notebooks := ResponseData{}
	service.GetNotebooks(notebooks)

	studios := service.GetStudios()
	pipelines := service.GetPipelines()

	var data []string
	data = append(data, notebooks.Data, studios, pipelines)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
