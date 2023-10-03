package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

// Возвращает подключение к базе данных
// Инициализирует базу
func NewMongodb(storagePath string, dbname string) (*mongodb, error) {
	const op = "storage.mongodb.NewMongodb"

	var ctx = context.TODO()
	opts := options.Client().ApplyURI(storagePath).SetTimeout(500 * time.Millisecond)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db := client.Database(dbname)

	//Готовим структуру и индексы
	//Для коллекции goods ставим ключевым поле gtin
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"gtin", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection("goods").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &mongodb{client: client, ctx: &ctx, db: db}, nil
}

// Закрыть подключение к базе данных
func (m *mongodb) Close() error {
	const op = "storage.mongodb.Close"
	err := m.client.Disconnect(*m.ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Добавляет продукт в хранилище
func (m *mongodb) AddGood(gtin string, description string) error {
	const op = "storage.mongodb.AddGood"

	g := Good{gtin, description}

	name, err := m.db.Collection("goods").InsertOne(*m.ctx, g)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Print(name)
	return nil
}
