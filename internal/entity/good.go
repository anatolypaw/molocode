package entity

import "time"

type Good struct {
	Gtin            Gtin
	Desc            string
	StoreCount      uint
	GetCodeForPrint bool
	AllowProduce    bool
	Upload          bool
	CreatedAt       time.Time
}
