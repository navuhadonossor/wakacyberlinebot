package main

import (
	"database/sql"
	"fmt"
	"os"
)

func openConnect() *sql.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName)
	db, _ := sql.Open("postgres", url)

	return db
}

func insertUser(db *sql.DB, tgId int, tgName string) {
	db.Exec(
		"INSERT INTO user (telegram_id, telegram_name) VALUES ($1, $2)",
		tgId,
		tgName,
	)
}

func updateUser(db *sql.DB, tgId int, apiToken string, wakatimeName string) {
	db.Exec(
		"UPDATE user SET api_token = $1 AND wakatime_name = $2 WHERE telegram_id = $3",
		apiToken,
		wakatimeName,
		tgId,
	)
}
