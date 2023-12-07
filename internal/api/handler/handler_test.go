package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler"
	"git.home/alex/go-subscriptions/tests"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		name           string
		handler        api_response.Handle
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "success",
			handler: func(_ *http.Request, _ httprouter.Params) any {
				return "success"
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"success","error":"","data":"success"}`,
		},
		{
			name: "error",
			handler: func(_ *http.Request, _ httprouter.Params) any {
				return tests.ErrTest
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"error","error":"` + tests.ErrTest.Error() + `","data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			handler.Handle(tc.handler)(w, nil, nil)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}
