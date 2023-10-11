package storage

import "fmt"

func (s *S) AddUser(login string, password string, role string) error {
	const op ="storage.AddUser"
	err := s.mongodb.AddUser(login, password, role)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}