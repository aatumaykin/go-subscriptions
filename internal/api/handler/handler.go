package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"github.com/julienschmidt/httprouter"
)

var errorMessage = `{"status":"error","error":"%s","data":null}`

func Handle(h api_response.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result := h(r, ps)

		if err, ok := result.(error); ok {
			writeError(w, err)

			return
		}

		if dto, ok := result.(api_response.ResponseDTO); ok {
			writeDTO(w, dto)

			return
		}

		dto := api_response.Success(result)
		writeDTO(w, dto)
	}
}

func write(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func writeError(w http.ResponseWriter, err error) {
	response := []byte(fmt.Sprintf(errorMessage, err.Error()))
	write(w, response)
}

func writeDTO(w http.ResponseWriter, dto api_response.ResponseDTO) {
	response, err := json.Marshal(dto)
	if err != nil {
		writeError(w, err)

		return
	}

	write(w, response)
}
