package mongostore

import "hub/internal/entity"

type FullCode_dto struct {
	Serial       string `bson:"_id"`
	Crypto       string
	SourceInfo   entity.SourceInfo     `bson:",omitempty"`
	PrintInfo    entity.PrintInfo      `bson:",omitempty"`
	ProducedInfo []entity.ProducedInfo `bson:",omitempty"`
	UploadInfo   entity.UploadInfo     `bson:",omitempty"`
}

type Counters struct {
	Name  string `bson:"_id"`
	Value uint64
}
