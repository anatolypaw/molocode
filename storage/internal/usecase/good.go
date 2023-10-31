package usecase

import (
	"fmt"
	"log"
	"storage/internal/adapter/storage"
	"storage/internal/entity"
)

// Добавляет продукт
func AddGood(s *storage.Connection, gtin string, desc string) error {
	const op = "usecase.AddGood"
	good, err := entity.New(gtin)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Установить описание
	good.EditDescription(desc)

	// TODO поместить в БД
	err = s.AddGood(gtin, desc)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("ОТЛАДКА продукт %v добавлен в БД", good)
	return nil
}

// Установить описание продукта
func EditDescription(s *storage.Connection, gtin string, desc string) error {
	const op = "usecase.EditDescription"
	good, err := entity.New(gtin)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	good.EditDescription(desc)

	// TODO поместить в БД
	log.Printf("ОТЛАДКА установлено описание продукта %v", good)
	return nil
}
