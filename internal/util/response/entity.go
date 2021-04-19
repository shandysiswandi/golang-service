package response

type SuccessBody struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorBody struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Reason  interface{} `json:"reason"`
}
