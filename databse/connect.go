package databse

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conect() {
	dsn := "host=localhost user=postgres password=salman dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
