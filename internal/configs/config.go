package configs

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env          string
	Port         int
	IsProduction bool

	Database  DatabaseConfig
	JWT       JWTConfig
	Cache     CacheConfig
	RateLimit RateLimitConfig
	CORS      CORSConfig
}

type DatabaseConfig struct {
	URI            string
	Database       string
	MaxPoolSize    uint64
	MinPoolSize    uint64
	SocketTimeout  time.Duration
	ConnectTimeout time.Duration
}

type JWTConfig struct {
	Secret    string
	ExpiresIn time.Duration
}

type CacheConfig struct {
	Size int
}

type RateLimitConfig struct {
	Window time.Duration
	Max    int
}

type CORSConfig struct {
	Origins []string
}

func Load() AppConfig {
	_ = godotenv.Load()

	env := getEnv("APP_ENV", "development")

	return AppConfig{
		Env:          env,
		Port:         getEnvInt("PORT", 4000),
		IsProduction: env == "production",

		Database: DatabaseConfig{
			URI:            getEnv("MONGODB_URI", "mongodb://localhost:27017"),
			Database:       getEnv("MONGODB_DB", "medcord"),
			MaxPoolSize:    100,
			MinPoolSize:    5,
			SocketTimeout:  45 * time.Second,
			ConnectTimeout: 5 * time.Second,
		},
		JWT: JWTConfig{
			Secret:    mustEnv("JWT_SECRET"),
			ExpiresIn: getEnvDuration("JWT_EXPIRES_IN", 7*24*time.Hour),
		},
		Cache: CacheConfig{
			Size: getEnvInt("CACHE_SIZE", 4096),
		},
		RateLimit: RateLimitConfig{
			Window: getEnvDuration("RATE_LIMIT_WINDOW", 15*time.Minute),
			Max:    getEnvInt("RATE_LIMIT_MAX", 100),
		},
		CORS: CORSConfig{
			Origins: splitAndTrim(getEnv("CORS_ORIGINS", "http://localhost:3000")),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("required env %s is missing", key)
	}
	return v
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}

func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}
