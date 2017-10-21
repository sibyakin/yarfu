package main

import (
	"github.com/gramework/gramework"
)

func main() {
	mux := gramework.New()
	mux.GET("/", "hello, grameworld")
	mux.ListenAndServe(":8080")
}
