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


/* Возвращает пользователя по его логину */
func (s *Storage) GetUser(login string) (structs.User, error) {
	const op ="storage.GetUser"
	u, err := s.mongodb.GetUser(login)
	if err != nil {
		return structs.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return u, err
}