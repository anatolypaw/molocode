package mongostore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION_GOODS    = "goods"
	COLLECTION_COUNTERS = "counters"
)

// TODO
type Cache struct {
	Goods struct {
		Value      []entity.Good
		LastUpdate time.Time
	}
}

type MongoStore struct {
	client *mongo.Client
	db     *mongo.Database
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*MongoStore, error) {
	const op = "mongo.New"
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
	}

	return &con, nil
}
