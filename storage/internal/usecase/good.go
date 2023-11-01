package usecase

import (
	"fmt"
	"storage/internal/adapter/storage"
	"storage/internal/entity"
)

// Добавляет продукт
func AddGood(storage *storage.Connection, good entity.Good) error {
	const op = "usecase.AddGood"
	err := good.ValidateGtin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// поместить в БД
	err = storage.AddGood(good)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Изменить настройки продукта
func EditGood(s *storage.Connection, good entity.Good) error {
	const op = "usecase.EditParams"
	err := good.ValidateGtin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// поместить в БД
	err = storage.EditGood(good)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
}
