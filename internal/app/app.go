package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"

	"github.com/spiderocious/medcord-backend/internal/configs"
	"github.com/spiderocious/medcord-backend/internal/deps"
	"github.com/spiderocious/medcord-backend/internal/middlewares"
	"github.com/spiderocious/medcord-backend/internal/routes"
	"github.com/spiderocious/medcord-backend/internal/shared/constants"
	"github.com/spiderocious/medcord-backend/internal/utils/response"
)

func New(cfg configs.AppConfig, d *deps.Dependencies) *gin.Engine {
	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.MaxMultipartMemory = 10 << 20 // 10 MB

	engine.Use(middlewares.Recovery(d.Logger, cfg.IsProduction))
	engine.Use(middlewares.RequestLogger(d.Logger))

	engine.Use(secure.New(secure.Config{
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		ContentSecurityPolicy: "default-src 'self'",
	}))

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.Origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Accept-Language"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.Use(middlewares.RateLimit(cfg.RateLimit.Max, cfg.RateLimit.Window))

	routes.Register(engine, d)

	engine.NoRoute(func(c *gin.Context) {
		response.NotFound(c, constants.MsgNotFound)
	})

	return engine
}
