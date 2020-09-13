package repository

import (
	"github.com/jinzhu/gorm"
	"go_rabbitmq_bet/models"
)

type LotteryBetRepository struct {

}

func NewLotteryBetRepository() *LotteryBetRepository{
	return &LotteryBetRepository{}
}

/**
	取得單一比注單
 */
func (repository *LotteryBetRepository) GetTotalUnsettleBet(lotteryId string, round string) (int, *gorm.DB){
	var count int
	query := models.Db.Model(&models.Bet{}).Where("lottery_id = ? AND round = ? AND status = 1",
		lotteryId, round).
		Order("created_at DESC").
		Count(&count)
	return count, query
}

