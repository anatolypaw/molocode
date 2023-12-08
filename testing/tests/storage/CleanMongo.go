package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CleanMongo(address string, db string) {

	fmt.Print("Очистка mongodb")
	//Подключаемся к хранилищу
	storage, err := New(address, db)
	if err != nil {
		log.Fatalln(err)
	}
	storage.Clean()
	storage.Close()
	fmt.Println(" == DONE ==")
}

type Storage struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

// Возвращает подключение к базе данных
func New(path string, dbname string) (*Storage, error) {
	const op = "storage.CleanMongo.New"

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

	con := Storage{
		client: client,
		ctx:    &ctx,
		db:     client.Database(dbname),
	}

	return &con, nil
}

// Удаляет базу
func (s *Storage) Clean() error {
	const op = "storage.ClenMongo.Clean"

	err := s.db.Drop(context.TODO())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Закрыть подключение к базе данных
func (con *Storage) Close() error {
	const op = "storage.CleanMongo.Close"
	err := con.client.Disconnect(*con.ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
