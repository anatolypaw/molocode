package usecase

import (
	"fmt"
	"log"
	"storage/internal/domain/entity"
)

// Добавляет продукт
func AddGood(gtin string, desc string) error {
	const op = "usecase.AddGood"
	good, err := entity.New(gtin)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Установить описание
	good.SetDescription(desc)

	// TODO поместить в БД
	log.Printf("ОТЛАДКА продукт %v добавлен в БД", good)
	return nil
}

// Установить описание продукта
func SetDescription(gtin string, desc string) error {
	const op = "usecase.SetDescription"
	good, err := entity.New(gtin)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	good.SetDescription(desc)

	// TODO поместить в БД
	log.Printf("ОТЛАДКА установлено описание продукта %v", good)
	return nil
}
