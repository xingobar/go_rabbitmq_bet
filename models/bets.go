package models

import "time"

type Bet struct {
	ID int `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserId int `gorm:"column:user_id" json:"user_id"`
	LotteryId int `gorm:"column:lottery_id" json:"lottery_id"`
	RoundId int `gorm:"column:round_id" json:"round_id"`
	BetAmount int `gorm:"column:bet_amount" json:"bet_amount"`
	ReturnAmount float64 `gorm:"column:return_amount" json:"return_amount"`
	WinAmount float64 `gorm:"column:win_amount" json:"win_amount"`
	Status int `gorm:"column:status" json:"status"`
	OrderId string `gorm:"column:order_id" json:"order_id"`
	BetDate string `gorm:"column:bet_date" json:"bet_date"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Title string `gorm:"column:title" json:"title"`
	Detail string `gorm:"column:detail" json:"detail"`
	OddsId int `gorm:"column:odds_id" json:"odds_id"`
	Odds float32 `gorm:"column:odds" json:"odds"`
	Round string `gorm:"column:round" json:"round"`
	Ball int `gorm:"column:ball" json:"ball"`
	Position int `gorm:"column:position" json:"position"`
	BetType int `gorm:"column:bet_type" json:"bet_type"`
	InfoId int `gorm:"column:info_id" json:"info_id"`
	UseRound int `gorm:"column:use_round" json:"use_round"`
	RealWin	float32 `gorm:"column:real_win" json:"real_win"`
}

func (Bet) TableName() string {
	return "lottery_bets"
}