package entity

import (
	"errors"
	"time"
)

type Good struct {
	Gtin            string
	Desc            string
	StoreCount      uint
	GetCodeForPrint bool
	AllowProduce    bool
	AllowPrint      bool
	Upload          bool
	CreatedAt       time.Time
}

func (ths *Good) ValidateDesc() error {
	if len(ths.Desc) < 3 || len(ths.Desc) > 40 {
		return errors.New("описание меньше 3 или длиннее 40 символов")
	}
	return nil
}

func (ths *Good) ValidateGtin() error {
	return ValidateGtin(ths.Gtin)
}
