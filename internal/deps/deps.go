package deps

import (
	"log/slog"

	"github.com/spiderocious/medcord-backend/internal/configs"
	"github.com/spiderocious/medcord-backend/internal/controllers"
	"github.com/spiderocious/medcord-backend/internal/utils/database"
)

// Dependencies holds every wired service, controller, and middleware used by routes.
// Built once at startup in Wire and passed into route registration.
type Dependencies struct {
	Logger *slog.Logger
	DB     *database.Mongo

	HealthController *controllers.HealthController
}

func Wire(_ configs.AppConfig, db *database.Mongo, logger *slog.Logger) *Dependencies {
	return &Dependencies{
		Logger:           logger,
		DB:               db,
		HealthController: controllers.NewHealthController(db),
	}
}
