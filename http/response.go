package http

import (
	"net/http"
)

type Response struct {
	Writer http.ResponseWriter
}

