package models

import "time"

type User struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Account string `gorm:"column:account" json:"account"`
	UsuableAmount float64 `gorm:"column:usuable_amount" json:"usuable_amount"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}