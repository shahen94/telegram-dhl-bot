package bot

import (
	"os"
	"telegram-dhl/pkg/bot/handlers"
	"telegram-dhl/pkg/database"
	"telegram-dhl/pkg/jobs"
	"telegram-dhl/pkg/logger"
	"telegram-dhl/pkg/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Repository *database.Repository
	Logger     *logger.Logger
	bot        *tgbotapi.BotAPI
	manager    *handlers.Manager
	jobs       *jobs.Jobs
}

func (b *Bot) Start() {
	b.Repository.Setup()
	b.Logger.Info("Authorized on account: ", b.bot.Self.UserName)

	services := services.NewOrderServices(b.Repository)

	b.manager = handlers.NewManager(b.bot, services)
	b.jobs = jobs.New(services)

	b.Logger.Info("Starting jobs...[Crawler]")
	b.jobs.Add(b.jobs.ScrapDHL)

	b.Logger.Info("Bot started")
	b.manager.Run()
}

func New() *Bot {
	apiKey := os.Getenv("TELEGRAM_API_KEY")

	if apiKey == "" {
		panic("TELEGRAM_API_KEY not set")
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)

	if err != nil {
		panic(err)
	}

	bot.Debug = true

	return &Bot{
		Repository: database.New(),
		Logger:     logger.New(),
		bot:        bot,
	}
}
