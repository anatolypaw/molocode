package mongodb

import (
	"fmt"
	"molocode/internal/structs"

	"go.mongodb.org/mongo-driver/bson"
)

//Добавляет пользователя в базу
func (m *Mongodb) AddUser(u structs.User) error {
	const op = "storage.mongodb.AddUser"

	_, err := m.db.Collection("users").InsertOne(*m.ctx, u)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

//Ищет пользователя по логину
func (m *Mongodb) GetUser(login string) (structs.User, error) {
	const op = "storage.mongodb.GetUser"

	filter := bson.D{{Key: "login", Value: login}}

	var u structs.User
	err := m.db.Collection("users").FindOne(*m.ctx, filter).Decode(&u)
	if err != nil {
		return structs.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return u, nil
}