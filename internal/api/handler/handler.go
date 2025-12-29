package handler

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	apiService "gitlab.com/english-vocab/telegram-bot/internal/api/service"
	"gitlab.com/english-vocab/telegram-bot/internal/telegram/bot"
	"gitlab.com/english-vocab/telegram-bot/internal/telegram/botService"
)

type Handler struct {
	services *apiService.Service
	tgBot    *bot.Bot
}

func NewHandler(services *apiService.Service, newBot *tgbotapi.BotAPI, service *botService.Service) *Handler {
	return &Handler{
		services: services,
		tgBot:    bot.NewBot(newBot, service),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	authorized := router.Group("/")

	authorized.Use(ValidateToken)
	{
		authorized.POST("/registration", h.setStatusRegistrationApplication)
		authorized.POST("/send-message", h.sendMessageToUsers)
	}

	return router
}
