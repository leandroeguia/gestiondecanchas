package handler

import (
	"fmt"
	"net/http"
)

// Esta es la única función que Vercel va a ejecutar
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Backend de canchas funcionando en Vercel!")
}
