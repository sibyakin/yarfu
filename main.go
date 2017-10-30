package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

const workdir = "./public/"

// String64 is a Gin request binding for json/base64 parsing
type String64 struct {
	Name    string `json:"name" binding:"required"`
	Image64 string `json:"image64" binding:"required"`
}

// URLForm here is a Gin request binding for url parsing
type URLForm struct {
	URLstring string `form:"url" binding:"required"`
}

func main() {
	mux := gin.Default()
	mux.Static("/public", "./public")

	// In future we can easily add new versions
	// of our API and It will be painless (hopefully)
	apiv1 := mux.Group("/api/v1")
	apiv1.POST("/images", imgAdd)
	apiv1.GET("/images/url", imgAddURL)
	apiv1.POST("/images/json", imgAddJSON)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, os.Interrupt)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server started: %s\n", err)
		}
	}()

	// This part will be executed only after receiving SIGINT
	<-gracefulStop
	log.Printf("Shutting down the server...\n")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server graceful shutdown error:", err)
	}
	log.Printf("Server exited")
}

func createThumb(filename string) {
	img, _ := imaging.Open(workdir + filename)
	thumb := imaging.Thumbnail(img, 100, 100, imaging.Box)
	result := imaging.New(100, 100, color.NRGBA{0, 0, 0, 0})
	result = imaging.Paste(result, thumb, image.Pt(0, 0))
	imaging.Save(result, workdir+"thumb_"+filename)
}

func imgAdd(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["files[]"]
	for _, file := range files {
		err := ctx.SaveUploadedFile(file, workdir+file.Filename)
		if err != nil {
			ctx.String(500, "Unable to save one or more file(s)!")
			return
		}
		createThumb(file.Filename)
	}
	ctx.String(200, fmt.Sprintf("%d file(s) uploaded successfully!\n", len(files)))
}

func imgAddJSON(ctx *gin.Context) {
	var json String64
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.String(500, "Unable to parse request body!\n")
		return
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(json.Image64))
	file, err := os.Create(workdir + json.Name)
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
	createThumb(json.Name)
	ctx.String(200, "Image saved successfully\n")
}

func imgAddURL(ctx *gin.Context) {
	var form URLForm
	err := ctx.Bind(&form)
	if err != nil {
		ctx.String(500, "Unable to parse request body!\n")
		return
	}
	tokens := strings.Split(form.URLstring, "/")
	filename := tokens[len(tokens)-1]
	file, err := os.Create(workdir + filename)
	if err != nil {
		ctx.String(500, "Unable to download requested file!\n")
		return
	}
	defer file.Close()
	response, err := http.Get(form.URLstring)
	if err != nil {
		ctx.String(500, "Unable to download requested file!\n")
		return
	}
	defer response.Body.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		ctx.String(500, "Unable to download requested file!\n")
		return
	}
	createThumb(filename)
	ctx.String(200, "Image saved successfully\n")
}
