package storage

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

/* Управление пользователями */
type User struct {
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"` //TODO передалать на хранение хэша
	Role     string `bson:"role" json:"role"`         //admin, user
}

// Инициализирует коллекцию users

// Добавляет пользователя в базу
func (m *Connection) AddUser(u User) error {
	const op = "storage.mongodb.AddUser"

	_, err := m.db.Collection("users").InsertOne(*m.ctx, u)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Ищет пользователя по логину
func (m *Connection) GetUser(login string) (User, error) {
	const op = "storage.mongodb.GetUser"

	filter := bson.D{{Key: "login", Value: login}}

	var u User
	err := m.db.Collection("users").FindOne(*m.ctx, filter).Decode(&u)
	if err != nil {
		return User{}, fmt.Errorf("%s: %w", op, err)
	}
	return u, nil
}

// Ищет пользователя по логину и паролю
func (m *Connection) GetUserByLoginPass(login string, password string) (User, error) {
	const op = "storage.mongodb.GetUserByLoginPass"

	filter := bson.D{{Key: "login", Value: login}, {Key: "password", Value: password}}

	var u User
	err := m.db.Collection("users").FindOne(*m.ctx, filter).Decode(&u)
	if err != nil {
		return User{}, fmt.Errorf("%s: %w", op, err)
	}
	return u, nil
}
