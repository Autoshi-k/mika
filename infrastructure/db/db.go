package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewDbConnection() Connection {
	username := "root"
	password := "root"
	hostAndPort := "127.0.0.1:3306"
	dbName := "mika"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, hostAndPort, dbName)

	conn, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed setting new db connection due to err " + err.Error())
	}

	return Connection{db: conn}
}
