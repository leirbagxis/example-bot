package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/leirbagxis/example-bot/internal/api"
	"github.com/leirbagxis/example-bot/internal/database"
	"github.com/leirbagxis/example-bot/internal/telegram"
)

// Send any text message to the bot after the bot has been started

func main() {
	db := database.InitDB()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	webhookHandler, botInstance := telegram.StartBot(db)

	go func() {
		if err := api.StartApi(db, webhookHandler, botInstance); err != nil {
			log.Printf("Erro ao iniciar API: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Println("🧹 Encerrando app com segurança...")

}
