package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/m3o"
	"github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/timeline"
)

const (
	MessageForbidden          = "You don't have the permission, please contact https://t.me/h1z3y3."
	MessageNotCommand         = "Please send a command. Send /help to get all commands"
	MessageUnsupportedCommand = "Unsupported command"
)

var (
	m3oApiKey        = os.Getenv("TG_M3O_TOKEN")
	telegramBotToken = os.Getenv("TG_BOT_TOKEN")

	adminList = map[string]struct{}{
		"h1z3y3":  {},
		"ankangz": {},
	}
)

// response 返回机器人回复的内容
func response(message *tgbotapi.Message) (string, error) {

	if _, ok := adminList[message.From.UserName]; !ok {
		return MessageForbidden, nil
	}

	if !message.IsCommand() {
		return MessageNotCommand, nil
	}

	db := m3o.NewDB(m3oApiKey).WithTablePrefix("tg_tm")
	tm := timeline.NewTimeline(db, time.Now().Format("200601"))

	var response string
	var err error

	switch message.Command() {
	case "add":
		response, err = tm.Add(timeline.Timeline{
			Content: message.CommandArguments(),
		})
	case "delete":
		response, err = tm.Delete(message.CommandArguments())
	case "read":
		response, err = tm.Read(message.CommandArguments())
	case "list_table":
		response, err = tm.ListTable()
	case "drop_table":
		response, err = tm.DropTable(message.CommandArguments())
	default:
		help := strings.Join([]string{
			"/add content: 添加时间线",
			"/delete id: 删除",
			"/read: 获取当前月份所有消息",
			"/list_table: 所有月份",
			"/drop_table 202206: 删除某个月份",
		}, "\n")
		response = MessageUnsupportedCommand + "\n\n" + help
	}

	if err != nil {
		response = err.Error()
	}

	return response, nil
}

// pollRobot 是轮询方式的机器人实现
func pollRobot() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		message := update.Message
		if message == nil { // ignore any non-Message Updates
			continue
		}

		resp, _ := response(update.Message)
		_, _ = bot.Send(tgbotapi.NewMessage(message.Chat.ID, resp))
	}
}

// webhookRobot 是 webhook 方式的机器人实现
func webhookRobot() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, err := tgbotapi.NewWebhook("https://telegram-timeline-robot.herokuapp.com/" + bot.Token)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	for update := range updates {
		message := update.Message
		if message == nil { // ignore any non-Message Updates
			continue
		}

		resp, _ := response(update.Message)
		_, _ = bot.Send(tgbotapi.NewMessage(message.Chat.ID, resp))
	}
}

func main() {
	//go pollRobot()
	go webhookRobot()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
		return
	})

	port := os.Getenv("PORT")

	fmt.Println("Listening on :" + port)

	http.ListenAndServe(":"+port, nil)
}
