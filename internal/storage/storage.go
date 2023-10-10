package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Интерфейсы
type Storage interface {

	//Закрывает подключение к хранилищу
	Close() error

	/* Управление продуктами файл goods.go*/
	//Выводит список продуктов в хранилище
	GetGoods() ([]Good, error)

	//Добавляет продукт в хранилище
	AddGood() error

	/* Управление пользователями файл users.go*/
	// Создать пользователя
	AddUser() error

}


// Код
type Code struct {
	Gtin   string // gtin продукта
	Serial string // серийный номер КМ
	Crypto string // криптохвост
}

type mongodb struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

// Возвращает подключение к базе данных
// Инициализирует базу
func NewMongodb(storagePath string, dbname string) (*mongodb, error) {
	const op = "storage.NewMongodb"

	var ctx = context.TODO()
	opts := options.Client().ApplyURI(storagePath).SetTimeout(500 * time.Millisecond)

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

	/* Готовим структуру и индексы */
	// Для коллекции goods ставим ключевым и уникальным поле gtin
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"gtin", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection("goods").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Коллеция users
	// ставим ключевым и уникальным поле login
	indexModel = mongo.IndexModel{
		Keys:    bson.D{{"login", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection("users").Indexes().CreateOne(ctx, indexModel)
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


