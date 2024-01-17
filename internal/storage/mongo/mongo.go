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

type MongoStore struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*MongoStore, error) {
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

	con := MongoStore{
		client: client,
		db:     client.Database(dbname),
		ctx:    context.TODO(),
	}

	return &con, nil
}
