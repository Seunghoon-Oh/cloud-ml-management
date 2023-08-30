package v1

import (
	"net/http"

	"github.com/Seunghoon-Oh/cloud-ml-manager/service"
	"github.com/gin-gonic/gin"
)

func GetAllMLServervices(c *gin.Context) {

	nobooks := service.GetNotebooks()
	studios := service.GetStudios()
	pipelines := service.GetPipelines()

	var data []string
	data = append(data, nobooks, studios, pipelines)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
