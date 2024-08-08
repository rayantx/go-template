package handlers

import (
	"net/http"

	"github.com/rayantx/gohtmxpr/views/foo"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	return foo.Index().Render(r.Context(), w)
}
