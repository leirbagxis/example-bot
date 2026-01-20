package callbacks

import (
	"github.com/go-telegram/bot"
	"github.com/leirbagxis/example-bot/internal/database"
	"github.com/leirbagxis/example-bot/internal/telegram/callbacks/about"
	"github.com/leirbagxis/example-bot/internal/telegram/callbacks/start"
)

func LoadCallbacksHandlers(b *bot.Bot, c *database.AppContainer) {
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "start", bot.MatchTypeExact, start.StartHandler(c))
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "about", bot.MatchTypeExact, about.AboutHandler(c))
}
