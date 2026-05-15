package main

import (
	"log"
	"os"

	"gestion-canchas/backend/internal/repository"

	"github.com/joho/godotenv" // Importas la librería
)

func main() {
	// Carga las variables desde el archivo .env
	// Si estás parado en la raíz, buscará el archivo .env ahí mismo
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: No se pudo cargar el archivo .env, se usarán las variables del sistema")
	}

	// Ahora sí, Go podrá leerla sin problemas
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("La variable de entorno DATABASE_URL no está configurada")
	}

	conn, err := repository.ConnectDB(connStr)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatalf("No se pudo responder al Ping de Supabase: %v", err)
	}

	log.Println("¡Conexión exitosa a Supabase desde el backend!")
}
