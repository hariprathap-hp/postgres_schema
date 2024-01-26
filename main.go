// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "altimetrik"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
	createTable(db)
	insertData(db, "altimetrik", "altimetrik@altimetrik.com")

}

func createTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50),
            email VARCHAR(100)
        );
    `

	_, err := db.Exec(query)
	return err
}

func insertData(db *sql.DB, username, email string) error {
	query := `
        INSERT INTO users (username, email) VALUES ($1, $2);
    `

	_, err := db.Exec(query, username, email)
	return err
}
