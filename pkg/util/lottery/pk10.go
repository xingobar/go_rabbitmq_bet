package lottery

import (
	"fmt"
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
	b, err := strconv.Atoi(pk10.result[bet.Position - 1])
	if err != nil {
		return false
	}
	ball := bet.Ball
	fmt.Println("postion: ", bet.Position, " ball: ", ball , " result: ", b, " bool: ", ball == b)
	return ball == b
}