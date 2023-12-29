package entity

type Counters struct {
	Name  string `bson:"_id"`
	Value uint32 `bson:"Value"`
}
