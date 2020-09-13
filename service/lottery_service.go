package service

import (
	"fmt"
	"go_rabbitmq_bet/models"
	"go_rabbitmq_bet/repository"
)

type LotteryService struct {
	betRepository *repository.LotteryBetRepository
}

func NewLotteryService() *LotteryService{
	return &LotteryService{
		betRepository: repository.NewLotteryBetRepository(),
	}
}

/**
	取得未結注單
	@param string lotteryId - 彩票編號
	@param string round - 期數
 */
func (s *LotteryService) GetTotalUnsettleBet(lotteryId string, round string) (int, error) {
	count, query := s.betRepository.GetTotalUnsettleBet(lotteryId, round)

	if err := query.Error; err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, fmt.Errorf("no unsettle data")
	}

	return count, nil
}

/**
	根據分頁取得注單
	@param string lotteryId - 彩票編號
	@param string round - 期數
	@param int page - 頁碼
 */
func (s *LotteryService) GetByPaginator(lotteryId string, round string, page int) ([]models.Bet, error){
	bets, query := s.betRepository.GetByPaginator(lotteryId, round, page)

	if err := query.Error; err != nil {
		return nil, err
	}
	return bets, nil
}
