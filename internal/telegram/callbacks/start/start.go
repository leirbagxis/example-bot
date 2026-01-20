package start

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/leirbagxis/example-bot/internal/database"
	"github.com/leirbagxis/example-bot/internal/utils"
	"github.com/leirbagxis/example-bot/pkg/parser"
)

func StartHandler(c *database.AppContainer) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		firstName := utils.RemoveHTMLTags(update.CallbackQuery.From.FirstName)

		text, button := parser.GetMessage("start", map[string]string{
			"firstName": firstName,
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
