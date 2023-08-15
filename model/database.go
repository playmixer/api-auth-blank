package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Cfg DSN
	DB  *gorm.DB
)

type DSN interface {
	Host() string
	Port() string
	Username() string
	Password() string
	DBName() string
}

func Init(cfg DSN) {
	Cfg = cfg
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&UserAuthCode{})
}

func Connect() (*gorm.DB, error) {
	var err error
	if DB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Cfg.Username(), Cfg.Password(), Cfg.Host(), Cfg.Port(), Cfg.DBName())
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	return DB, err
}
