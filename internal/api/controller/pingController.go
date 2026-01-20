package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leirbagxis/example-bot/internal/database"
)

func PingHandler(c *database.AppContainer) gin.HandlerFunc {
	return func(g *gin.Context) {
		res := map[string]any{
			"ping": "pong",
		}
		g.JSON(http.StatusOK, res)
	}
}
