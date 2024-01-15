package hubstorage

import "time"

type Good_dto struct {
	Gtin            string `bson:"_id"`
	StoreCount      uint
	GetCodeForPrint bool
	AllowProduce    bool
	Upload          bool
	CreatedAt       time.Time
}
