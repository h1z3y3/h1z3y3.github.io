package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/httplib"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/go-github/v45/github"
	dbIface "github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/db"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	"github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/m3o"
	"github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/timeline"
)

const (
	MessageForbidden          = "You don't have the permission, please contact https://t.me/h1z3y3."
	MessageNotCommand         = "Please send a command. Send /help to get all commands"
	MessageUnsupportedCommand = "Unsupported command"
	MessageError              = "ERROR: "
)

var (
	m3oApiKey        = os.Getenv("TG_M3O_TOKEN")
	telegramBotToken = os.Getenv("TG_BOT_TOKEN")

	githubOwner         = os.Getenv("TG_GITHUB_OWNER")
	githubRepo          = os.Getenv("TG_GITHUB_REPO")
	githubToken         = os.Getenv("TG_GITHUB_TOKEN")
	githubWorkflowId, _ = strconv.Atoi(os.Getenv("TG_GITHUB_WORKFLOW_ID"))

	adminList = map[string]struct{}{
		"h1z3y3":  {},
		"ankangz": {},
	}
)

var (
	db  dbIface.DB
	tm  timeline.Interface
	bot *tgbotapi.BotAPI
	gh  *github.Client
)

func githubActionsUrl() string {
	return fmt.Sprintf("https://github.com/%s/%s/actions/runs/%d", githubOwner, githubRepo, githubWorkflowId)
}

func helpMessage() string {
	help := strings.Join([]string{
		"/add content: 添加时间线",
		"/delete id: 删除",
		"/read: 获取当前月份所有消息",
		"/last: 获取当前月份最新消息",
		"/list_table: 所有月份",
		"/drop_table 202206: 删除某个月份",
	}, "\n")

	return help
}

// responsePhoto 回复 图片消息
func uploadPhotoToGithub(photo tgbotapi.PhotoSize) (string, error) {
	// get from telegram
	f, err := bot.GetFile(tgbotapi.FileConfig{FileID: photo.FileID})
	if err != nil {
		return "", errors.Wrap(err, "get file by fileId error: "+photo.FileID)
	}

	fp := "images/" + f.FilePath
	branch := "emo-update"

	// if already on the github
	ghFile, _, _, _ := gh.Repositories.GetContents(context.Background(),
		githubOwner, githubRepo, fp,
		&github.RepositoryContentGetOptions{Ref: branch})

	if ghFile != nil {
		return ghFile.GetDownloadURL(), nil
	}

	// get file content
	bs, err := httplib.Get(f.Link(telegramBotToken)).
		SetTimeout(1*time.Second, 10*time.Second).
		Retries(3).RetryDelay(1 * time.Second).
		Bytes()

	if err != nil {
		return "", errors.Wrap(err, "get file content error")
	}

	// upload to github
	cfg := &github.RepositoryContentFileOptions{
		Message: github.String("image auto commit: " + photo.FileID),
		Content: bs,
		Branch:  github.String(branch),
	}

	resp, _, err := gh.Repositories.CreateFile(context.Background(), githubOwner, githubRepo, fp, cfg)
	if err != nil {
		return "", errors.Wrap(err, "upload to github error")
	}

	return resp.Content.GetDownloadURL(), nil
}

// responseText 回复 文本消息
func responseCommand(message *tgbotapi.Message) (string, error) {

	var response string
	var err error

	switch message.Command() {
	case "add":
		response, err = tm.Add(timeline.Timeline{
			Content: message.CommandArguments(),
		})

		if err != nil {
			break
		}
		_, err = gh.Actions.RerunWorkflowByID(context.Background(), githubOwner, githubRepo, int64(githubWorkflowId))

		response += "\n\n Triggered Github Action: " + githubActionsUrl()
	case "delete":
		response, err = tm.Delete(message.CommandArguments())
		if err != nil {
			break
		}
		_, err = gh.Actions.RerunWorkflowByID(context.Background(), githubOwner, githubRepo, int64(githubWorkflowId))
		response += "\n\n Triggered Github Action: " + githubActionsUrl()
	case "read":
		var list timeline.Timelines
		list, err = tm.Read(message.CommandArguments())
		response = list.String()
	case "last":
		var last timeline.Timeline
		last, err = tm.LastTimeline("")
		response = last.String()
	case "list_table":
		response, err = tm.ListTable()
	case "drop_table":
		response, err = tm.DropTable(message.CommandArguments())
	default:
		response = MessageUnsupportedCommand + "\n\n" + helpMessage()
	}

	if err != nil {
		response = err.Error()
	}

	return response, nil
}

// response 返回机器人回复的内容
func response(message *tgbotapi.Message) {
	chatId := message.Chat.ID
	if _, ok := adminList[message.From.UserName]; !ok {
		_, _ = bot.Send(tgbotapi.NewMessage(chatId, MessageForbidden))
		return
	}

	// command
	if message.IsCommand() {
		resp, err := responseCommand(message)
		if err != nil {
			_, _ = bot.Send(tgbotapi.NewMessage(chatId, MessageError+err.Error()))
			return
		}

		_, _ = bot.Send(tgbotapi.NewMessage(chatId, resp))
		return
	}

	// photo
	if len(message.Photo) > 0 {
		photos := message.Photo
		maxSizePhoto := photos[len(photos)-1]

		downloadUrl, err := uploadPhotoToGithub(maxSizePhoto)
		if err != nil {
			_, _ = bot.Send(tgbotapi.NewMessage(chatId, MessageError+err.Error()))
			return
		}

		// last
		last, err := tm.LastTimeline("")
		if err != nil {
			_, _ = bot.Send(tgbotapi.NewMessage(chatId, MessageError+err.Error()))
			return
		}

		last.Images = append(last.Images, downloadUrl)
		err = tm.Update(last)
		if err != nil {
			_, _ = bot.Send(tgbotapi.NewMessage(chatId, MessageError+err.Error()))
			return
		}

		_, _ = bot.Send(tgbotapi.NewMessage(chatId, last.String()))
		return
	}

	_, _ = bot.Send(tgbotapi.NewMessage(chatId, "帮助信息:\n\n"+helpMessage()))

	return
}

// pollRobot 是轮询方式的机器人实现
func pollRobot() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		log.Fatal(err)
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

		response(update.Message)
	}
}

// webhookRobot 是 webhook 方式的机器人实现
func webhookRobot() {
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

		response(update.Message)
	}
}

func init() {
	var err error

	//telegram bot
	bot, err = tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// m3o storage
	db = m3o.NewDB(m3oApiKey).WithTablePrefix("tg_tm")
	tm = timeline.NewTimeline(db, time.Now().Format("200601"))

	// github client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh = github.NewClient(tc)
}

func main() {
	env := os.Getenv("ENV")
	fmt.Println("env: ", env)
	if env == "" {
		fmt.Println("env is empty, use prod")
	}
	if env == "dev" {
		go pollRobot()
	} else {
		go webhookRobot()
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
		return
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on :" + port)

	http.ListenAndServe(":"+port, nil)
}
