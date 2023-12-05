package error_handler

import (
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
)

func HandleError(w http.ResponseWriter, err error) {
	responseDto := api_response.ErrorResponse
	responseDto.Error = err.Error()
	response, err := responseDto.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}
