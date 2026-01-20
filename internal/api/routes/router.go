package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leirbagxis/example-bot/internal/api/controller"
	"github.com/leirbagxis/example-bot/internal/database"
)

func RegisterRoutes(r *gin.Engine, c *database.AppContainer) {
	api := r.Group("/api")
	api.Use()
	{
		api.GET("/ping", controller.PingHandler(c))
	}
}
