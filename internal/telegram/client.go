package telegram

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/leirbagxis/example-bot/internal/cache"
	"github.com/leirbagxis/example-bot/internal/database"
	"github.com/leirbagxis/example-bot/internal/telegram/callbacks"
	"github.com/leirbagxis/example-bot/internal/telegram/commands"
	"github.com/leirbagxis/example-bot/internal/telegram/middleware"
	"github.com/leirbagxis/example-bot/pkg/config"
	"gorm.io/gorm"
)

func StartBot(db *gorm.DB) (http.Handler, *bot.Bot) {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	opts := []bot.Option{
		bot.WithMiddlewares(
			middleware.SaveUserMiddleware(db),
		),
	}

	b, err := bot.New(config.TelegramBotToken, opts...)
	if err != nil {
		panic(err)
	}

	cache.GetRedisClient()
	app := database.NewAppContainer(db, b)

	botInfo, _ := b.GetMe(ctx)
	botUsername := fmt.Sprintf("@%s", botInfo.Username)
	log.Println("🤖 Bot iniciado:", botUsername)

	originalHandler := b.WebhookHandler()
	debugHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("❌ Erro ao ler body: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		originalHandler.ServeHTTP(w, r)
		log.Println("✅ Webhook processado com sucesso")
	})

	go func() {
		<-ctx.Done()
		log.Println("🔻 Shutting down gracefully...")
		if err := cache.CloseRedis(); err != nil {
			log.Printf("❌ Error closing Redis: %v", err)
		}
		cancel()
	}()

	webhookUrl := config.WebhookURL
	if webhookUrl != "" {
		log.Printf("🔗 Bot configurado para modo webhook: %s", webhookUrl)

		commands.LoadCommandHandlers(b, app)
		callbacks.LoadCallbacksHandlers(b, app)

		_, err := b.SetWebhook(ctx, &bot.SetWebhookParams{
			URL: webhookUrl,
			//AllowedUpdates: []string{"message", "callback_query", "inline_query", "my_chat_member"},
		})
		if err != nil {
			log.Fatalf("❌ Erro ao setar webhook: %v", err)
		}

		log.Println("✅ Webhook configurado com sucesso")

		webhookInfo, err := b.GetWebhookInfo(ctx)
		if err == nil {
			log.Printf("📊 Webhook Info - URL: %s, Pending: %d",
				webhookInfo.URL, webhookInfo.PendingUpdateCount)
		}

		log.Println("🚀 Iniciando webhook...")
		go b.StartWebhook(ctx)

	} else {
		log.Println("🔄 Bot iniciado em modo polling")
		commands.LoadCommandHandlers(b, app)
		callbacks.LoadCallbacksHandlers(b, app)

		b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{})
		go b.Start(ctx)
	}

	return debugHandler, b
}
