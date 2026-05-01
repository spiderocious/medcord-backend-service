package database

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/spiderocious/medcord-backend/internal/configs"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect(ctx context.Context, cfg configs.DatabaseConfig, logger *slog.Logger) (*Mongo, error) {
	opts := options.Client().
		ApplyURI(cfg.URI).
		SetMaxPoolSize(cfg.MaxPoolSize).
		SetMinPoolSize(cfg.MinPoolSize).
		SetSocketTimeout(cfg.SocketTimeout).
		SetConnectTimeout(cfg.ConnectTimeout)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	logger.Info("mongo connected", "db", cfg.Database)
	return &Mongo{Client: client, DB: client.Database(cfg.Database)}, nil
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

func (m *Mongo) Ping(ctx context.Context) error {
	return m.Client.Ping(ctx, readpref.Primary())
}
