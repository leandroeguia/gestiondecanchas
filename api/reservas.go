package api

import (
	"fmt"
	"net/http"
)

func ReservaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint de reservas funcionando correctamente")
}
