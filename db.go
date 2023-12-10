package main

import (
	"database/sql"
	"log"
	"os"
)

func openConnect() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Open connection failed: " + err.Error())
	}

	return db
}

func insertUser(db *sql.DB, tgId int, tgName string) {
	_, err := db.Exec(
		"INSERT INTO user (telegram_id, telegram_name) VALUES ($1, $2)",
		tgId,
		tgName,
	)
	if err != nil {
		log.Println("Insert user failed: " + err.Error())
	}
}

func updateUser(db *sql.DB, tgId int, apiToken string, wakatimeName string) {
	_, err := db.Exec(
		"UPDATE user SET api_token = $1 AND wakatime_name = $2 WHERE telegram_id = $3",
		apiToken,
		wakatimeName,
		tgId,
	)
	if err != nil {
		log.Println("Update user failed: " + err.Error())
	}
}
