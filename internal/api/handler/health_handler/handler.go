package health_handler

import (
	"net/http"

	"git.home/alex/go-subscriptions/internal/version"
	"github.com/julienschmidt/httprouter"
)

func Handle() httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		data := `{"status":"pass","version":"` + version.Version + `"}`

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(data))
	}
}
