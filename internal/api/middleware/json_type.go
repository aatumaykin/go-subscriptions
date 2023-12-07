package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetJSONContentType(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r, ps)
	}
}
