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
	Id               string `json:"id"`
	TelegramId       int    `json:"telegram_id"`
	TelegramName     string `json:"telegram_name"`
	WakatimeApiToken string `json:"wakatime_api_token"`
	WakatimeName     string `json:"wakatime_name"`
}
