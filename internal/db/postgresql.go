package db

import (
	"os"

	"github.com/nano2nano/valorant_tips/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db := getSession()
	db.AutoMigrate(&model.Agent{}, &model.Ability{}, &model.Map{}, &model.Tip{}, &model.Side{})
	return db
}

func getSession() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		return db
	}
}
