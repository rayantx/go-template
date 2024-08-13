package handlers

import (
	"net/http"

	"github.com/rayantx/gohtmxpr/views/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
