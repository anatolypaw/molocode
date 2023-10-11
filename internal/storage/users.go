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
func (s *Storage) CheckUserPassword(login string, password string) (bool, error) {
	const op ="storage.ChekUserPassword"
	u, err := s.mongodb.GetUser(login)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return "pass", err
}