package main

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)

var (
	bot     *tgbotapi.BotAPI
	baseURL = "https://wakacyberlinebot-269a92218149.herokuapp.com/"
)

func initTelegram() {
	var err error

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Println(err)
		return
	}

	url := baseURL + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "HELLO")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	initTelegram()
	router.POST("/"+bot.Token, webhookHandler)

	err := router.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
