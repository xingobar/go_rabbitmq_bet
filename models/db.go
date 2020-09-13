package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	args := fmt.Sprintf("%s:@/%s?parseTime=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_DATABASE"))
	Db, err = gorm.Open("mysql", args)
	if err != nil {
		panic("error load mysql")
	}
}