package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	mux := gin.Default
	mux.GET("/api/v1/images", imageslist)
	mux.Run()
}

func imageslist() {

}
