package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-manager/controller/v1"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/ml/services", v1.GetAllMLServervices)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8082")
}
