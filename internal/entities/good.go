package entities

import "time"

type Good struct {
	Gtin            *Gtin
	StoreCount      uint
	GetCodeForPrint bool
	AllowProduce    bool
	Upload          bool
	CreatedAt       time.Time
}
