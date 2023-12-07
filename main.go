package wakatime_tg_bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"os"
	"reflect"
	"time"
)

func main() {
	time.Sleep(1 * time.Minute)
	runBot()
}

func runBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "qweqweqwe")
				bot.Send(msg)
			}
		}
	}
}
