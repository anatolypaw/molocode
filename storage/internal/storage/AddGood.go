package storage

import (
	"fmt"
	"log"
	"regexp"
)

// Добавляет продукт в хранилище
func (s *Storage) AddGood(gtin string, desc string) error {
	const op = "storage.AddGood"

	//Проверяем корректность gtin
	matched, err := regexp.MatchString(`^0\d{13}$`, gtin)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		log.Print(err)
		return err
	}

	if !matched {
		return fmt.Errorf("%s: Некорректный формат gtin %s", op, gtin)
	}


	g := Good{
		Gtin: gtin,
		Desc: desc,
	}

	_, err = s.db.Collection("goods").InsertOne(*s.ctx, g)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}