package hubstore

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionGoods    = "goods"
	collectionCounters = "counters"
)

type hubStorage struct {
	client *mongo.Client
	db     *mongo.Database
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*hubStorage, error) {
	const op = "hubstorage.NewHubStorage"
	opts := options.Client().ApplyURI(path).SetTimeout(1000 * time.Millisecond)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Проверка подключения к базе
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	con := hubStorage{
		client: client,
		db:     client.Database(dbname),
	}

	return &con, nil
}
