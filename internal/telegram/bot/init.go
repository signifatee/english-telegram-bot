package bot

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/english-vocab/telegram-bot/pkg"
)

func Init(envPath string) {
	if err := godotenv.Load(envPath); err != nil {
		logrus.Fatal("No .env file found")
	}

	err, msg := pkg.ValidateConfigs()
	if err != true {
		logrus.Fatal(msg)
	}
}
