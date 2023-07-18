package config

import (
	"fmt"
	"log"

	"com.ktb.ai.core-item-catalog/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	TimeZone string
}

func NewDBInstance(config *DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v", config.Host, config.Username, config.Password, config.Database, config.Port, config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db err ", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(100)
	return db
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.ItemCatalog{})
}
