package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Probar la conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
