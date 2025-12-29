package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/internal/api/handler"
	"gitlab.com/english-vocab/telegram-bot/internal/api/server"
	apiService "gitlab.com/english-vocab/telegram-bot/internal/api/service"
	"gitlab.com/english-vocab/telegram-bot/internal/share/repository"
	"gitlab.com/english-vocab/telegram-bot/internal/telegram/bot"
	"gitlab.com/english-vocab/telegram-bot/internal/telegram/botService"
	"os"
)

// rand.Seed(time.Now().UnixNano())

func main() {
	//body := strings.NewReader("{\n    \"message\": \"Тест Animals был обновлен, пройдите его заново\",\n    \"users\": [\n        {\n            \"externalId\": \"13255443654\"\n        },\n        {\n            \"externalId\": \"13255463564\"\n        },\n        {\n            \"externalId\": \"63546563654\"\n        },\n        {\n            \"externalId\": \"31546578654\"\n        }\n    ]\n}")
	//var send dto.SendMessageToUsers
	//json.NewDecoder(body).Decode(&send)
	//
	//logrus.Infof(send.Message)
	//for _, user := range send.Users {
	//	logrus.Infof(user.ChatId)
	//}

	bot.Init("./.env")

	tgBot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		panic(err)
	}

	tgBot.Debug = true

	db_cfg := repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	postgres, err := repository.NewPostgresDB(db_cfg)
	if err != nil {
		logrus.Fatalf("Cannot open connection to PostgresDB: %s", err)
	}
	repos := repository.NewRepository(postgres)
	botSvc := botService.NewService(repos)

	srv_cfg := server.NewConfig()
	services := apiService.NewService(repos)
	handlers := handler.NewHandler(services, tgBot, botSvc)
	srv := server.New(srv_cfg, handlers.InitRoutes())

	telegramBot := bot.NewBot(tgBot, botSvc)

	go func() {
		if err := srv.Start(); err != nil {
			logrus.Fatalf("Cannot start api due to: \n %s", err.Error())
		}
	}()

	go func() {
		if err := telegramBot.Start(); err != nil {
			logrus.Fatal(err)
		}
	}()
	select {}
}
