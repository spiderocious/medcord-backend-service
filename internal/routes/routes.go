package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/spiderocious/medcord-backend/internal/deps"
)

func Register(engine *gin.Engine, d *deps.Dependencies) {
	api := engine.Group("/api")
	api.GET("/health", d.HealthController.Check)
}
