package api_response

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handle func(*http.Request, httprouter.Params) any
