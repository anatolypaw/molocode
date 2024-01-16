package hubservice

import (
	"fmt"
	"molocode/internal/domain/entity"
	"time"
)

// Добавляет код
func (hs *hubService) AddCodeForPrint(code entity.Code, sourceName string) error {
	err := code.Validate()
	if err != nil {
		return err
	}

	// Проверяем, существует ли такой продукт
	good, err := hs.storage.GetGood(code.Gtin)
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

	err = hs.storage.AddCode(mappedCode)
	if err != nil {
		return err
	}

	return nil
}

func (hs *hubService) GetCodeForPrint(gtin entity.Gtin, sourceName string) error {
	return nil
}
