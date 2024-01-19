package mongostore

import (
	"molocode/internal/app/entity"
	"time"
)

type Good_dto struct {
	Gtin            string `bson:"_id"`
	Desc            string
	StoreCount      uint
	GetCodeForPrint bool
	AllowProduce    bool
	Upload          bool
	CreatedAt       time.Time
}

type Code_dto struct {
	Serial       string `bson:"_id"`
	Crypto       string
	SourceInfo   entity.SourceInfo
	PrintInfo    entity.PrintInfo
	ProducedInfo []entity.ProducedInfo
	UploadInfo   entity.UploadInfo
}
