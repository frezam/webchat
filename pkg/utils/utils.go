package utils

import "encoding/json"

type Response struct {
	Error   bool        `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetInternalError() []byte {
	data, _ := json.Marshal(Response{
		Error:   true,
		Code:    500,
		Message: "Internal server error",
		Data:    nil,
	})
	return data
}

func (r *Response) Marshall() []byte {
	data, errMarshall := json.Marshal(*r)
	if errMarshall != nil {
		return GetInternalError()
	}
	return data
}
