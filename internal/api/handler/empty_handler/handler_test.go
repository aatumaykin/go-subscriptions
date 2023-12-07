package empty_handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/empty_handler"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		name     string
		expected api_response.ResponseDTO
	}{
		{
			name:     "Not Implemented Error",
			expected: api_response.Error(empty_handler.ErrNotImplemented),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			empty_handler.Handle()(w, nil, nil)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			expectedBody, err := json.Marshal(tc.expected)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
