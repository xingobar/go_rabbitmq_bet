package settle_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_rabbitmq_bet/pkg/e"
	"go_rabbitmq_bet/resource"
	"go_rabbitmq_bet/service"
)

type SettleController struct {
	lotteryService *service.LotteryService
}

func NewSettleController() *SettleController {
	return &SettleController{
		lotteryService: service.NewLotteryService(),
	}
}

func (c *SettleController) Settle(ctx *gin.Context) {
	lotteryId := ctx.Param("id")
	round := ctx.Param("round")
	ball := ctx.Param("ball")
	fmt.Println("id: ", lotteryId)
	fmt.Println("round: ", round)
	fmt.Println("ball: ", ball)

	bet, err := c.lotteryService.FetchByFilter(lotteryId, round)
	if err != nil {
		resource.ErrorResponse(ctx, e.NOT_FOUND)
		return
	}
	fmt.Println(bet)
}