package health_handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	"git.home/alex/go-subscriptions/internal/api/handler/health_handler"
	"git.home/alex/go-subscriptions/internal/version"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
	}{
		{
			name:     "success",
			expected: `{"status":"pass","version":"` + version.Version + `"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := &http.Request{}
			ps := httprouter.Params{}

			health_handler.Handle()(w, r, ps)

			assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "handler returned wrong content type")
			assert.Equal(t, tc.expected, w.Body.String(), "handler returned unexpected body")
		})
	}
}
