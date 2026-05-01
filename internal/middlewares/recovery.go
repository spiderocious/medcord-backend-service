package middlewares

import (
	"log/slog"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/spiderocious/medcord-backend/internal/shared/constants"
	"github.com/spiderocious/medcord-backend/internal/utils/response"
)

func Recovery(logger *slog.Logger, isProduction bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				fields := []any{
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
					"panic", rec,
				}
				if !isProduction {
					fields = append(fields, "stack", string(debug.Stack()))
				}
				logger.Error("panic recovered", fields...)
				response.ServerError(c, constants.MsgInternalServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
