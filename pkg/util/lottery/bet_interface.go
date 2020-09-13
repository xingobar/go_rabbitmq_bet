package lottery

import "go_rabbitmq_bet/models"

type BetInterface interface {
	Settle(bet models.Bet) bool // 結算
}