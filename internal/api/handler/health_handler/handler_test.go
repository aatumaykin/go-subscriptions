package health_handler_test

import (
	"context"
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
			// Create a new router and add the route
			router := httprouter.New()
			router.GET("/health", health_handler.Handle())

			// Create a new request
			req, err := http.NewRequestWithContext(context.Background(), "GET", "/health", nil)
			assert.NoError(t, err)

			// Create a new ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			// Call the router's ServeHTTP method to execute the request
			router.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			assert.Equal(t, tc.expectedStatus, rr.Code, "handler returned wrong status code")
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "handler returned wrong content type")

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), rr.Body.String(), "handler returned unexpected body")
		})
	}
}
