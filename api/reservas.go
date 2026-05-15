package handler

import (
	"fmt"
	"net/http"
)

// Esta función DEBE llamarse Handler y ser pública (Mayúscula)
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "¡El backend de Gestión de Canchas está vivo!")
}
