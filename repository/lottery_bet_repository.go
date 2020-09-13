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
	@param lotteryId - 彩票編號
	@param round - 期數
 */
func (repository *LotteryBetRepository) GetTotalUnsettleBet(lotteryId string, round string) (int, *gorm.DB){
	var count int
	query := models.Db.Model(&models.Bet{}).Where("lottery_id = ? AND round = ? AND status = 1",
		lotteryId, round).
		Order("created_at DESC").
		Count(&count)
	return count, query
}

/**
	根據頁碼去得注單
	@param int lotteryId - 彩票編號
	@param string round - 期數
	@param int page - 頁碼
 */
func (repository *LotteryBetRepository) GetByPaginator(lotteryId string, round string, page int) ([]models.Bet, *gorm.DB){
	var bets []models.Bet
	query := models.Db.
		Where("lottery_id = ? AND round = ? AND status = 1", lotteryId, round).
		Find(&bets)
	return bets, query
}
