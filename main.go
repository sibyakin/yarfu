package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mux := gin.Default()
	mux.LoadHTMLGlob("templates/*.html")
	mux.Static("/static", "./static")

	apiv1 := mux.Group("/api/v1")
	apiv1.GET("/images", imageslist)
	apiv1.GET("/images/:id", imagesdetail)
	apiv1.PUT("/images", imagesadd)

	log.Fatalln(mux.Run())
}

func imageslist(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func imagesdetail(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func imagesadd(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
