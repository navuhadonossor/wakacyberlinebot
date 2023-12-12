package main

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"os"
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
		client := &http.Client{}
		url := os.Getenv("WAKATIME_HOST") + "/api/compat/wakatime/v1/users/current"
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Failed to construct request!")
		}
		request.Header.Add("Authorization", apiToken)
		response, err := client.Do(request)
		if err != nil {
			log.Println("Response failed: " + err.Error())
		}
		var jsonResponse wakatimeUserResponse
		bytes, _ := io.ReadAll(response.Body)
		log.Println("WAKA RESPONSE")
		log.Println(string(bytes))
		_ = json.Unmarshal(bytes, &jsonResponse)
		updateUser(from.ID, apiToken, jsonResponse.data.id)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Now you join leaderboard :smirk: Good luck!")
		msg.ReplyToMessageID = message.MessageID
		bot.Send(msg)
	}
}

func generateMembersList(message *tgbotapi.Message) {
	text := "Members: \n"
	users, _ := getUserList()
	for i, user := range users {
		row := strconv.Itoa(i + 1)
		text = text + row + ". " + user.TelegramName + "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}

func generateLaderboardTable(message *tgbotapi.Message) {
	generateMembersList(message)
}
