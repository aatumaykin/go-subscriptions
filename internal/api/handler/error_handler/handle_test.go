package error_handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/tests"
	"github.com/stretchr/testify/assert"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
)

func TestHandleError(t *testing.T) {
	testCases := []struct {
		name           string
		inputError     error
		expectedStatus int
		expectedBody   api_response.ResponseDTO
	}{
		{
			name:           "error",
			inputError:     tests.ErrTest,
			expectedStatus: http.StatusOK,
			expectedBody: api_response.ResponseDTO{
				Status: "error",
				Error:  tests.ErrTest.Error(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := http.NewRequest("GET", "/error", nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()

			error_handler.HandleError(w, tc.inputError)

			assert.Equal(t, tc.expectedStatus, w.Code, "handler returned wrong status code")
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "handler returned wrong content type")

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String(), "handler returned unexpected body")
		})
	}
}
