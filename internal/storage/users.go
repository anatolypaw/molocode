package storage

import "fmt"

/* Управление пользователями */
type User struct {
	Login string 	`bson:"login" json:"login"`
	Password string `bson:"password" json:"password"` //TODO передалать на хранение хэша
	Role string 	`bson:"role" json:"role"`	  //admin, user
}


//Добавляет пользователя в базу
func (m *mongodb) AddUser(login string, password string, role string) error {
	const op = "storage.AddUser"

	u := User {Login: login, Password: password, Role: role}

	_, err := m.db.Collection("users").InsertOne(*m.ctx, u)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}