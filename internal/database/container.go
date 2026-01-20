package database

import (
	"github.com/go-telegram/bot"
	"github.com/leirbagxis/example-bot/internal/cache"
	"github.com/leirbagxis/example-bot/internal/database/repository"
	"github.com/leirbagxis/example-bot/internal/database/service"
	"gorm.io/gorm"
)

type AppContainer struct {
	DB          *gorm.DB
	Bot         *bot.Bot
	UserService *service.UserService

	// ## CACHE ## \\
	CacheService *cache.Service
}

func NewAppContainer(db *gorm.DB, bot *bot.Bot) *AppContainer {
	cacheService := cache.NewService()
	//tele := teleclient.NewTelegramAdapter(bot)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	return &AppContainer{
		DB:           db,
		UserService:  userService,
		Bot:          bot,
		CacheService: cacheService,
	}
}
