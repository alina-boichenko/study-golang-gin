package main

import (
	"github.com/alina-boichenko/study-golang-gin.git/controller"
	"github.com/alina-boichenko/study-golang-gin.git/service"
	"github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	videoService service.VideoService = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

// @title Study Golang Gin
// @version 1.0
// @description This is a sample server
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	server := gin.Default()
	// add swagger
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
