package health_handler_test

import (
	"encoding/json"
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
		name           string
		expectedStatus int
		expectedBody   health_handler.ResponseDTO
	}{
		{
			name:           "success",
			expectedStatus: http.StatusOK,
			expectedBody: health_handler.ResponseDTO{
				Status:  "pass",
				Version: version.Version,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := &http.Request{}
			ps := httprouter.Params{}

			health_handler.Handle()(w, r, ps)

			assert.Equal(t, tc.expectedStatus, w.Code, "handler returned wrong status code")
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "handler returned wrong content type")

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String(), "handler returned unexpected body")
		})
	}
}
