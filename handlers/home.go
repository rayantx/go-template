package handlers

import (
	"net/http"

	"github.com/rayantx/gohtmxpr/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
