package settle_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_rabbitmq_bet/pkg/e"
	"go_rabbitmq_bet/pkg/rabbitmq"
	"go_rabbitmq_bet/pkg/setting"
	"go_rabbitmq_bet/resource"
	"go_rabbitmq_bet/service"
	"math"
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

	// 彩種找不到
	// TODO: 搬移到 middleware
	if _, ok := setting.LotteryQueueMap[lotteryId]; !ok {
		resource.ErrorResponse(ctx, e.NOT_FOUND)
		return
	}

	// 先取得未結注單
	count, err := c.lotteryService.GetTotalUnsettleBet(lotteryId, round)
	if err != nil {
		resource.ErrorResponse(ctx, e.NOT_FOUND)
		return
	}

	total := int(math.Ceil(float64(count) / setting.PER_PAGE))
	fmt.Printf("total data %d \n", count)
	fmt.Printf("total page %d \n", total)

	queueName := setting.GetQueueNameByLottery(lotteryId)
	for page := 1; page <= total; page++ {
		bets, err := c.lotteryService.GetByPaginator(lotteryId, round, page)
		if err != nil {
			resource.ErrorResponse(ctx, e.ERROR)
			break
		}
		// publish message to queue
		fmt.Println("publish message ....")
		rabbitmq.Publish(queueName, bets)

		// consume message
		fmt.Println("consume message ....")
		rabbitmq.ConsumeBets(queueName)
	}

	fmt.Println(count)
}