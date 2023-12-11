package main

const JsonFilepath = "users.json"

type wakatimeUserResponse struct {
	data wakatimeUserData
}
type wakatimeUserData struct {
	id       string
	username string
}

type User struct {
	id               string `json:"id"`
	telegramId       int    `json:"telegram_id"`
	telegramName     string `json:"telegram_name"`
	wakatimeApiToken string `json:"wakatime_api_token"`
	wakatimeName     string `json:"wakatime_name"`
}
