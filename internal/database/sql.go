package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"rwa/config"
)

type Sql struct {
	DB *sql.DB
}

func InitSqlDB() (*Sql, error) {
	cfg := config.LoadConfig()
	var username = cfg.MySql.Username
	var password = cfg.MySql.Password
	var host = cfg.MySql.Host
	var port = cfg.MySql.Port
	var dbName = cfg.MySql.Database

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/?parseTime=true")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("USE " + dbName)
	if err != nil {
		return nil, err
	}

	return &Sql{DB: db}, nil
}
