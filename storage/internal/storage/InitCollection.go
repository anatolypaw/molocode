package storage

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin       string // gtin продукта
	Desc       string // описание продукта
	StoreCount int    // сколько хранить кодов
	Get        bool   // флаг, получать коды из 1с
	Upload     bool   // флаг, выгружать коды в 1с
	Awaible    bool   // флаг, выдавать ли кода на терминал
	ShelfLife  int    // срок годности продукта
}


// Инициализирует коллекцию goods
func (s *Storage)InitCollectionGoods() error {
	const op = "storage.goodsInitCollection"

	// Для коллекции goods ставим ключевым и уникальным поле gtin
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"gtin", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := s.db.Collection("goods").Indexes().CreateOne(*s.ctx, indexModel)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}


