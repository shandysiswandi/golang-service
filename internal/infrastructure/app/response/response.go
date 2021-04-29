package response

import (
	"encoding/json"
)

type (
	SuccessBody struct {
		Error   bool        `json:"error"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	ErrorBody struct {
		Error   bool        `json:"error"`
		Message string      `json:"message"`
		Reason  interface{} `json:"reason"`
	}
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

func Error(msg string, reason interface{}) ErrorBody {
	return ErrorBody{Error: true, Message: msg, Reason: reason}
}

func BadRequest() {
	//
}
