package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/spiderocious/medcord-backend/internal/shared/constants"
	"github.com/spiderocious/medcord-backend/internal/utils/database"
	"github.com/spiderocious/medcord-backend/internal/utils/response"
)

type HealthController struct {
	db    *database.Mongo
	start time.Time
}

func NewHealthController(db *database.Mongo) *HealthController {
	return &HealthController{db: db, start: time.Now()}
}

type HealthStatus struct {
	Status   string `json:"status"`
	DB       string `json:"db"`
	UptimeMS int64  `json:"uptimeMs"`
}

func (h *HealthController) Check(c *gin.Context) {
	pingCtx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "ok"
	if err := h.db.Ping(pingCtx); err != nil {
		dbStatus = "down"
	}

	response.Success(c, HealthStatus{
		Status:   "ok",
		DB:       dbStatus,
		UptimeMS: time.Since(h.start).Milliseconds(),
	}, constants.MsgHealthOK)
}
