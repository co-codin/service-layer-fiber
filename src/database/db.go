package database

import (
	"github/co-codin/service-layer-fiber/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(
		postgres.Open("postgres://postgres:root@localhost:5432/service_layer_fiber"),
		&gorm.Config{},
	)

	if err != nil {
		panic("can not connect to postgres")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
