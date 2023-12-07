package wakatime_tg_bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"net/http"
	"os"
	"reflect"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		runBot()
	})
}

func runBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/")
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
