package storeservice

import (
	"fmt"
	"molocode/internal/entity"
	"time"
)

// Добавляет код
func (hs *Service) AddCodeForPrint(code entity.Code, sourceName string) error {
	err := code.Validate()
	if err != nil {
		return err
	}

	// Проверяем, существует ли такой продукт
	good, err := hs.store.GetGood(code.Gtin)
	if err != nil {
		return err
	}

	// Проверяем, разрешено ли получение кодов для печати
	if !good.GetCodeForPrint {
		return fmt.Errorf("для продукта %s запрещено получение кодов для печати", code.Gtin)
	}

	// TODO проверять, не превышено ли нужное количество кодов

	// MAPPING
	mappedCode := entity.FullCode{
		Code: code,
		SourceInfo: entity.SourceInfo{
			Name: sourceName,
			Time: time.Now(),
		},
	}

	err = hs.store.AddCode(mappedCode)
	if err != nil {
		return err
	}

	return nil
}

func (hs *Service) GetCodeForPrint(gtin string, sourceName string) error {
	return nil
}
