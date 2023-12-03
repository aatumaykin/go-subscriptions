package handler

import (
	"net/http"

	"git.home/alex/go-subscriptions/internal/version"
	"github.com/julienschmidt/httprouter"
)

func Health() httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		_, _ = w.Write([]byte(`{"status":"pass", "version":"` + version.Version + `"}`))
	}
}
