package main

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

var (
	bot *tgbotapi.BotAPI
)

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

func initTelegram() {
	var err error

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Println(err)
		return
	}

	url := os.Getenv("HOST") + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := io.ReadAll(c.Request.Body)
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
	log.Print("INCOMING MESSAGE")
	marshaled, _ := json.Marshal(update)
	log.Println(string(marshaled))
	if update.Message == nil {
		return
	}
	switch update.Message.Text {
	case "/register":
		registerUser(update.Message)
	case "/top_today", "/top_today@wakacyberline_bot":
		generateLaderboardTable(update.Message, "today")
	case "/top_week", "/top_week@wakacyberline_bot":
		generateLaderboardTable(update.Message, "last_7_days")
	case "/top_month", "/top_month@wakacyberline_bot":
		generateLaderboardTable(update.Message, "last_30_days")
	case "/top_year", "/top_year@wakacyberline_bot":
		generateLaderboardTable(update.Message, "year")
	default:
		updateUserWakaToken(update.Message, update.Message.Text)
	}
	return
}
