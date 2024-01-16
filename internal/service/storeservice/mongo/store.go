package mongo

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

type Store struct {
	client *mongo.Client
	db     *mongo.Database
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*Store, error) {
	const op = "hubstorage.NewHubStore"
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

	con := Store{
		client: client,
		db:     client.Database(dbname),
	}

	return &con, nil
}
