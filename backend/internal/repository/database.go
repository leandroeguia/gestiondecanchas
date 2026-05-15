// backend/internal/repository/database.go
package repository

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// ConnectDB se escribe con C mayúscula para que sea accesible desde main.go
func ConnectDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	// Verificar si la conexión es real
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
