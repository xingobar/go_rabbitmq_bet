package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go_rabbitmq_bet/controller/settle_controller"
)

func main() {

	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		panic("load env file error")
	}

	settleController := settle_controller.NewSettleController()
	r.GET("/settle/:id/:round/:ball", settleController.Settle)

	r.Run()
}