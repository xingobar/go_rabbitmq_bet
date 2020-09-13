package service

import (
	"fmt"
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
