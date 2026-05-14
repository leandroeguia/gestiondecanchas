package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Definimos el endpoint que Nuxt buscará vía Proxy
	http.HandleFunc("/api/hola", func(w http.ResponseWriter, r *http.Request) {
		// Importante para que el front entienda que es JSON
		w.Header().Set("Content-Type", "application/json")

		data := map[string]string{
			"mensaje": "¡Conexión exitosa! Go y Nuxt están hablando.",
			"status":  "OK",
		}

		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("🚀 Servidor de Go corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
