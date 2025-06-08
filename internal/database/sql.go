package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"rwa/config"
)

type Sql struct {
	db *sql.DB
}

func InitSqlDB() (*Sql, error) {
	sqlConfig := config.Databases["sql"]
	var username = sqlConfig["username"]
	var password = sqlConfig["password"]
	var dbPort = sqlConfig["port"]
	var dbHost = sqlConfig["host"]
	var dbName = sqlConfig["dbname"]

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/")
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

	return &Sql{db: db}, nil
}

func (s *Sql) GetDB() *sql.DB {
	return s.db
}
