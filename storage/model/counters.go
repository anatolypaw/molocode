package model

type Counters struct {
	Name  string `bson:"_id" json:",omitempty"`
	Value uint32 `bson:"" json:",omitempty"`
}
