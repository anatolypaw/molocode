package storage

import (
	"fmt"
	"molocode/internal/structs"
)

/* Добавляет пользователя в базу */
func (s *Storage) AddUser(u structs.User) error {
	const op ="storage.AddUser"
	err := s.mongodb.AddUser(u)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}


/* Возвращает пользователя */
func (s *Storage) GetUser(login string) (structs.User, error) {
	const op = "storage.GetUser"
	// Запрашиваем пользователя из базы
	u, err := s.mongodb.GetUser(login)
	if err != nil {
		return structs.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return u, err
}

/* Возвращает пользователя, по паре логин и пароль */
func (s *Storage) GetUserByLoginPass(login string, password string) (structs.User, error) {
	const op ="storage.GetUserByLoginPass"

	// Запрашиваем пользователя из базы
	u, err := s.mongodb.GetUserByLoginPass(login, password)
	if err != nil {
		return structs.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return u, nil
}