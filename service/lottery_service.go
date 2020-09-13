package service

import (
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

func (s *LotteryService) FetchByFilter(lotteryId string, round string) (*models.Bet, error) {
	bet, query := s.betRepository.FetchByFilter(lotteryId, round)

	if err := query.Error; err != nil {
		return nil, err
	}
	return &bet, nil
}
