package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mux := gin.Default()
	mux.LoadHTMLGlob("templates/*.html")
	mux.Static("/public", "./public")

	// In future we can easily add new versions
	// of our API and It will be painless (hopefully)
	apiv1 := mux.Group("/api/v1")
	apiv1.GET("/images", imgListV1)
	apiv1.GET("/images/:id", imgDetailV1)
	apiv1.POST("/images", imgAddV1)
	apiv1.POST("/images/json", imgAddJSONV1)
	apiv1.POST("/images/url", imgAddURLV1)

	log.Fatalln(mux.Run())
}

func imgListV1(ctx *gin.Context) {
	ctx.HTML(200, "list.html", gin.H{})
}

func imgDetailV1(ctx *gin.Context) {
	ctx.HTML(200, "detail.html", gin.H{})
}

func imgAddV1(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		err := ctx.SaveUploadedFile(file, "./public/"+file.Filename)
		if err != nil {
			ctx.String(500, "Internal Server Error")
		}
	}
	ctx.String(200, fmt.Sprintf("%d file(s) uploaded successfully!\n", len(files)))
}

func imgAddJSONV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func imgAddURLV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
