package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-restful-api/helper"
	"go-restful-api/utils"
	"time"
)

func NewDB() *sqlx.DB {
	config, _ := utils.LoadConfig()
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBDatabase))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
