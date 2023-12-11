package main

type wakatimeUserResponse struct {
	data wakatimeUserData
}
type wakatimeUserData struct {
	id       string
	username string
}

type User struct {
	id               string
	telegramId       int
	telegramName     string
	wakatimeApiToken string
	wakatimeName     string
}
