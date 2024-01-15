package hubstorage

import (
	"context"
	"fmt"
	"molocode/internal/domain/entity"
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
func NewHubStorage(path string, dbname string) (*hubStorage, error) {
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

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (hs *hubStorage) AddGood(g entity.Good) error {
	const op = "hubstorage.AddGood"
	// MAPPING
	mappedGood := Good_dto{
		Gtin:            string(g.Gtin),
		StoreCount:      g.StoreCount,
		GetCodeForPrint: g.GetCodeForPrint,
		AllowProduce:    g.AllowProduce,
		Upload:          g.Upload,
		CreatedAt:       g.CreatedAt,
	}
	_, err := hs.db.Collection(collectionGoods).InsertOne(context.TODO(), mappedGood)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
