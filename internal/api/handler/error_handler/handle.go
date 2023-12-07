package error_handler

import (
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
)

func HandleError(w http.ResponseWriter, err error) {
	response, err := api_response.Error(err).ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}
