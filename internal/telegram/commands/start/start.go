package start

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/leirbagxis/example-bot/internal/utils"
	"github.com/leirbagxis/example-bot/pkg/parser"
)

func StartHandler() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		firstName := utils.RemoveHTMLTags(update.Message.From.FirstName)

		text, button := parser.GetMessage("start", map[string]string{
			"firstName": firstName,
		})

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        text,
			ReplyMarkup: button,
			ParseMode:   "HTML",
			ReplyParameters: &models.ReplyParameters{
				MessageID: update.Message.ID,
			},
		})
		if err != nil {
			fmt.Println("Erro enviar /start : ", err)
		}
	}
}
