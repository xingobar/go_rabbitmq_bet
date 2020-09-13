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
func (repository *LotteryBetRepository) FetchByFilter(lotteryId string, round string) (models.Bet, *gorm.DB){
	var Bets models.Bet
	query := models.Db.Where("lottery_id = ? AND round = ?", lotteryId, round).First(&Bets)
	return Bets, query
}

