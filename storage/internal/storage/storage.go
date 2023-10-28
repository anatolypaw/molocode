package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*Mongo, error) {
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

	db := client.Database(dbname)
	return &Mongo{client: client, ctx: &ctx, db: db}, nil
}

/* Инициализирует структуру базы */
func (m *Mongo) InitCollection() error {
	const op = "storage.mongodb.InitCollection"

	/* Готовим структуру и индексы */
	// Для коллекции goods ставим ключевым и уникальным поле gtin
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"gtin", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := m.db.Collection("goods").Indexes().CreateOne(*m.ctx, indexModel)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Коллеция users
	// ставим ключевым и уникальным поле login
	indexModel = mongo.IndexModel{
		Keys:    bson.D{{"login", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = m.db.Collection("users").Indexes().CreateOne(*m.ctx, indexModel)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Закрыть подключение к базе данных
func (m *Mongo) Close() error {
	const op = "storage.mongodb.Close"
	err := m.client.Disconnect(*m.ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}


