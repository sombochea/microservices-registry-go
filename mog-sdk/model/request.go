package model

import "encoding/json"

type SendDefinition struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (t *SendDefinition) ToJson() string {
	jsonBytes, _ := json.Marshal(t)
	return string(jsonBytes)
}

func FromJson(valueJson string) (*SendDefinition, error) {
	object := &SendDefinition{}
	err := json.Unmarshal([]byte(valueJson), object)
	return object, err
}
