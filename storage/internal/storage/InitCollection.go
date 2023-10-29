package storage

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Инициализирует коллекцию goods
func (s *Connection)InitCollectionGoods() error {
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


