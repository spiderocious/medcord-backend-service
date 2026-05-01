package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memstore "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimit(max int, window time.Duration) gin.HandlerFunc {
	rate := limiter.Rate{Period: window, Limit: int64(max)}
	store := memstore.NewStore()
	instance := limiter.New(store, rate)
	return ginlimiter.NewMiddleware(instance)
}

func AuthRateLimit() gin.HandlerFunc {
	return RateLimit(5, 15*time.Minute)
}
