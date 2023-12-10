package main

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func registerUser(message *tgbotapi.Message) {
	from := message.From
	insertUser(openConnect(), from.ID, from.UserName)
	msg := tgbotapi.NewMessage(message.Chat.ID, "Please send your wakatime API token in next message")
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
		bytes, _ := ioutil.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &jsonResponse)
		updateUser(openConnect(), from.ID, apiToken, jsonResponse.data.id)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Now you join leaderboard :smirk: Good luck!")
		msg.ReplyToMessageID = message.MessageID
		bot.Send(msg)
	}
}

func generateMembersList() {

}

func generateLaderboardTable() {

}
