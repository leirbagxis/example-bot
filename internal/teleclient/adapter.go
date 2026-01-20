package teleclient

import (
	"context"

	tg "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TelegramAdapter struct {
	bot *tg.Bot
}

func NewTelegramAdapter(bot *tg.Bot) *TelegramAdapter {
	return &TelegramAdapter{bot: bot}
}

func (a *TelegramAdapter) CreateInvoiceLink(ctx context.Context, params *tg.CreateInvoiceLinkParams) (string, error) {
	return a.bot.CreateInvoiceLink(ctx, params)
}

func (a *TelegramAdapter) SendMessage(ctx context.Context, params *tg.SendMessageParams) (*models.Message, error) {
	return a.bot.SendMessage(ctx, params)
}

func (a *TelegramAdapter) GetUsernameBOT(ctx context.Context, params *tg.SendMessageParams) string {
	botInfo, _ := a.bot.GetMe(ctx)
	return "@" + botInfo.Username
}
