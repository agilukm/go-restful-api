package app

import (
	"github.com/jmoiron/sqlx"
	"go-restful-api/helper"
	"go-restful-api/utils"
	"time"
)

func NewDB() *sqlx.DB {
	config, err := utils.LoadConfig()
	db, err := sqlx.Open("mysql", config.DBUsername+":"+config.DBPassword+"@tcp("+config.DBHost+":"+config.DBPort+")/"+config.DBDatabase)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
