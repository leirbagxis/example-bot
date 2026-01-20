package commands

import (
	"github.com/go-telegram/bot"
	"github.com/leirbagxis/example-bot/internal/database"
	"github.com/leirbagxis/example-bot/internal/telegram/commands/start"
)

func LoadCommandHandlers(b *bot.Bot, c *database.AppContainer) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, start.StartHandler())

}
