package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mux := gin.Default()
	mux.LoadHTMLGlob("templates/*.html")
	mux.Static("/static", "./static")

	// In future we can easily add new versions
	// of our API and It will be painless (hopefully)
	apiv1 := mux.Group("/api/v1")
	apiv1.GET("/images", imagesListV1)
	apiv1.GET("/images/:id", imageDetailV1)
	apiv1.PUT("/images", imageAddV1)

	log.Fatalln(mux.Run())
}

func imagesListV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func imageDetailV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func imageAddV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
