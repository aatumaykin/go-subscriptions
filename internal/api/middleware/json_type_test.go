package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestSetJSONContentType(t *testing.T) {
	handler := httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Do nothing
	})
	wrappedHandler := middleware.SetJSONContentType(handler)

	w := httptest.NewRecorder()
	r := &http.Request{}
	ps := httprouter.Params{}

	wrappedHandler(w, r, ps)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
