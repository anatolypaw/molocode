package mongostore

import (
	"molocode/internal/app/entity"
)

type Code_dto struct {
	Serial       string `bson:"_id"`
	Crypto       string
	SourceInfo   entity.SourceInfo
	PrintInfo    entity.PrintInfo
	ProducedInfo []entity.ProducedInfo
	UploadInfo   entity.UploadInfo
}

type Counters struct {
	Name  string `bson:"_id"`
	Value uint64
}
