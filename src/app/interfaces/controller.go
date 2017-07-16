package interfaces

import (
	"net/http"
)

type Controller interface {
	Handle(response http.ResponseWriter, request *http.Request)
}
