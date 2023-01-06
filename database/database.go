package database

import (
	"fmt"
	"log"
	"poc-go/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlRepository struct {
	connectionPool *gorm.DB
}

func InitDatabase(config utils.MysqlConfig) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@/poc?parseTime=true&loc=Local", config.User, config.Password)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	return db, err
}

func NewMysqlRepository() *mysqlRepository {
	config := utils.GetConfig().MysqlConfig

	dns := fmt.Sprintf("%s:%s@/poc?parseTime=true", config.User, config.Password)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return &mysqlRepository{
		connectionPool: db,
	}
}
