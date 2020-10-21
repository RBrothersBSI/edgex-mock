package main

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Pong)
	r.POST("v1/seedDevices/:count", handlers.Device)
	r.GET("/v1/device", handlers.GetAll)
	r.Run(":48082")
}
