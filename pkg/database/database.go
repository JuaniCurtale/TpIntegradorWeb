package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConectarDB() *sql.DB { //CONEXION A LA BASE DATOS
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	fmt.Println("--- LEYENDO VARIABLES DE ENTORNO ---")
	fmt.Printf("DB_HOST: [%s]\n", host)
	fmt.Printf("DB_PORT: [%s]\n", port)
	fmt.Printf("DB_USER: [%s]\n", user)
	fmt.Printf("DB_NAME: [%s]\n", dbname)
	fmt.Println("------------------------------------")

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
