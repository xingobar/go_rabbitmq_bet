package lottery

import (
	"go_rabbitmq_bet/models"
	"strconv"
)

type Pk10 struct {
	result []string
	bets models.Bet
}

func NewPk10(result []string) BetInterface{
	return &Pk10{result:result}
}

func (pk10 *Pk10) Settle(bet models.Bet) bool {
	betType := bet.BetType
	var win bool
	switch betType {
		case 1: {
			win = pk10.OneToTen(bet)
		}
	}
	return win
}

// 單球
func (pk10 *Pk10) OneToTen(bet models.Bet) bool{
	b := pk10.result[bet.Position]
	ball, _ := strconv.Atoi(b)
	return ball == bet.Ball
}