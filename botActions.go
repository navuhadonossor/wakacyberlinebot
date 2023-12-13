package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"sort"
	"strconv"
)

func registerUser(message *tgbotapi.Message) {
	err := insertUser(message.From.ID, message.From.UserName)
	text := "Please send your wakatime API token in next message"
	if err != nil {
		text = "You already registered!"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}

func updateUserWakaToken(message *tgbotapi.Message, apiToken string) {
	if message.Chat.IsPrivate() {
		from := message.From
		jsonResponse := getWakaUserInfo(apiToken)
		updateUser(from.ID, apiToken, jsonResponse.Data.Id)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Now you join leaderboard 😏 Good luck!")
		msg.ReplyToMessageID = message.MessageID
		bot.Send(msg)
	}
}

func generateLaderboardTable(message *tgbotapi.Message, period string) {
	text := "Leaderboard: \n"
	users, _ := getUserList()
	var responses []wakatimeStatResponse
	for _, user := range users {
		responses = append(responses, getUserStatInfo(user.WakatimeApiToken, period))
	}
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Data.Total_Seconds > responses[i].Data.Total_Seconds
	})
	for i, response := range responses {
		row := strconv.Itoa(i + 1)
		seconds := response.Data.Total_Seconds
		hours := strconv.Itoa(seconds / 60 / 60)
		minutes := strconv.Itoa((seconds % (60 * 60)) / 60)
		text = text + row + ". " + response.Data.Username + " " + hours + "h. " + minutes + "m.\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}
