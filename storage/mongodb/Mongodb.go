package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionGoods = "goods"
)

type Storage struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*Storage, error) {
	const op = "storage.mongodb.New"

	var ctx = context.TODO()
	opts := options.Client().ApplyURI(path).SetTimeout(100 * time.Millisecond)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Проверка подключения к базе
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	con := Storage{
		client: client,
		ctx:    &ctx,
		db:     client.Database(dbname),
	}

	// Инициализация коллекций
	// err = con.InitCollectionGoods()
	// if err != nil {
	//	return nil, fmt.Errorf("%s: %w", op, err)
	// }

	return &con, nil
}

// Закрыть подключение к базе данных
func (con *Storage) Close() error {
	const op = "storage.mongodb.Close"
	err := con.client.Disconnect(*con.ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Инициализирует коллекцию goods
func (s *Storage) InitCollectionGoods() error {
	const op = "storage.goodsInitCollection"

	// Для коллекции goods ставим ключевым и уникальным поле gtin
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "gtin", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := s.db.Collection("goods").Indexes().CreateOne(*s.ctx, indexModel)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
