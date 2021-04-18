package response

import (
	"encoding/json"
)

func Success(msg string, data interface{}) interface{} {
	return SuccessBody{Error: false, Message: msg, Data: data}
}

func SuccessForTest(body string) (*SuccessBody, error) {
	data := &SuccessBody{}
	if err := json.Unmarshal([]byte(body), data); err != nil {
		return nil, err
	}
	return data, nil
}
