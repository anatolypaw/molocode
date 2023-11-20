package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	goodCollection = "goods"
)

type Connection struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*Connection, error) {
	const op = "storage.New"

	var ctx = context.TODO()
	opts := options.Client().ApplyURI(path).SetTimeout(500 * time.Millisecond)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Проверка подключения к базе
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	con := Connection{
		client: client,
		ctx:    &ctx,
		db:     client.Database(dbname),
	}

	// Инициализация коллекций
	err = con.InitCollectionGoods()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &con, nil
}

// Закрыть подключение к базе данных
func (con *Connection) Close() error {
	const op = "storage.mongodb.Close"
	err := con.client.Disconnect(*con.ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
