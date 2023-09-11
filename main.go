package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-manager/controller/v1"
	"github.com/Seunghoon-Oh/cloud-ml-manager/service"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/ml/services", v1.GetAllMLServervices)
	r.POST("/ml/notebook", v1.CreateMLNotebook)
	r.POST("/ml/studio", v1.CreateMLStudio)
	r.POST("/ml/pipeline", v1.CreateMLPipeline)

	return r
}

func setupCircuitBreaker() {
	service.SetupNotebookCircuitBreaker()
	service.SetupPipelineCircuitBreaker()
	service.SetupStudioCircuitBreaker()
}

func main() {
	setupCircuitBreaker()
	r := setupRouter()
	r.Run(":8082")
}
