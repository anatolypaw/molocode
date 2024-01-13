package models

import "encoding/json"

type Response struct {
	Ok   bool   `json:"ok"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func (r *Response) ToJson() string {
	resultJson, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(resultJson)
}
