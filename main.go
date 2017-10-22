package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// String64 is a Gin request binding for json/base64 parsing
type String64 struct {
	Name    string `json:"name" binding:"required"`
	Image64 string `json:"image64" binding:"required"`
}

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
	files := form.File["files[]"]
	for _, file := range files {
		err := ctx.SaveUploadedFile(file, "./public/"+file.Filename)
		if err != nil {
			ctx.String(500, "Internal Server Error")
		}
	}
	ctx.String(200, fmt.Sprintf("%d file(s) uploaded successfully!\n", len(files)))
}

func imgAddJSONV1(ctx *gin.Context) {
	var json String64
	err := ctx.BindJSON(&json)
	if err != nil {
		fmt.Println(err)
		ctx.String(500, "Unable to parse request body!\n")
		return
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(json.Image64))
	file, err := os.Create("./public/" + json.Name)
	if err != nil {
		ctx.String(500, "Unable to save file on server!\n")
		return
	}
	defer file.Close()
	_, err = io.Copy(file, reader)
	if err != nil {
		ctx.String(500, "Unable to save file on server!\n")
		return
	}
	ctx.String(200, "Image saved successfully\n")
}

func imgAddURLV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
