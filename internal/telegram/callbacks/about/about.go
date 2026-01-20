package about

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/leirbagxis/example-bot/internal/database"
	"github.com/leirbagxis/example-bot/pkg/config"
	"github.com/leirbagxis/example-bot/pkg/parser"
)

func AboutHandler(c *database.AppContainer) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		botInfo, _ := b.GetMe(ctx)
		user, _ := b.GetChat(ctx, &bot.GetChatParams{
			ChatID: config.OwnerID,
		})

		text, button := parser.GetMessage("about", map[string]string{
			"ownerUser":  user.FirstName,
			"botVersion": "beta 1.0.8",
			"botId":      fmt.Sprintf("%d", botInfo.ID),
		})

		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			Text:        text,
			ReplyMarkup: button,
			ParseMode:   "HTML",
			MessageID:   update.CallbackQuery.Message.Message.ID,
		})
	}
}
