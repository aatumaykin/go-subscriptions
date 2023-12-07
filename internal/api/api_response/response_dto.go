package api_response

import "encoding/json"

type ResponseDTO struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func Success(data any) ResponseDTO {
	return ResponseDTO{Status: "success", Data: data}
}

func Error(err error) ResponseDTO {
	return ResponseDTO{Status: "error", Error: err.Error()}
}

func (r ResponseDTO) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}
