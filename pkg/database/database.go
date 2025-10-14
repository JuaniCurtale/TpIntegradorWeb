package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectarDB() *sql.DB { //CONEXION A LA BASE DATOS
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la conexión: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la DB exitosa")
	return dbConn
}
