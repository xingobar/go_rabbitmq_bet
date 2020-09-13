package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		panic("load env file error")
	}

	r.Run()
}