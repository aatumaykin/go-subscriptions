package api_response

import "encoding/json"

type ResponseDTO struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

var (
	SuccessResponse = ResponseDTO{Status: "success"}
	ErrorResponse   = ResponseDTO{Status: "error"}
)

func (r ResponseDTO) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}
