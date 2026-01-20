package teleclient

import (
	"context"

	tg "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TelegramClient interface {
	CreateInvoiceLink(ctx context.Context, params *tg.CreateInvoiceLinkParams) (string, error)
	SendMessage(ctx context.Context, params *tg.SendMessageParams) (*models.Message, error)
	GetUsernameBOT(ctx context.Context) string
}
