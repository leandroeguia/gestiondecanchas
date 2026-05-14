package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola desde Go! Estás en el backend de las canchas.")
}
