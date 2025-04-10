package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	dbName = "order_management"
)

func main() {

	var err error
	db, err = connectToDB(dbName)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/reserve", reserveDeliveryAgent)

	http.HandleFunc("/assign", assignDeliveryAgent)

	http.ListenAndServe(":8081", nil)
}

func reserveDeliveryAgent(w http.ResponseWriter, r *http.Request) {
	// Get the first available delivery agent and mark reserved

	txn, err := db.Begin()
	if err != nil {
		panic(err)
	}

	row := txn.QueryRow(`SELECT id
	FROM delivery_reservations
	WHERE is_reserved = false and 
	order_id IS null
	FOR UPDATE SKIP LOCKED
	LIMIT 1;`)

	var id int
	if err := row.Scan(&id); err != nil {
		fmt.Println("could not get row id")
		panic(err)
	}

	if _, err = txn.Exec(`UPDATE delivery_reservations SET is_reserved=true WHERE id = $1`, id); err != nil {
		panic(err)
	}

	if err := txn.Commit(); err != nil {
		panic(err)
	}

	fmt.Println("received request to reserve delivery agent")
}

func assignDeliveryAgent(w http.ResponseWriter, r *http.Request) {

	order_id := "order_id"

	txn, err := db.Begin()
	if err != nil {
		panic(err)
	}

	row := txn.QueryRow(`SELECT id
	FROM delivery_reservations
	WHERE is_reserved = true and 
	order_id IS null
	FOR UPDATE SKIP LOCKED
	LIMIT 1;`)

	var id int
	if err := row.Scan(&id); err != nil {
		panic(err)
	}

	if _, err = txn.Exec(`UPDATE delivery_reservations SET order_id=$1 WHERE id = $2`, order_id, id); err != nil {
		panic(err)
	}

	if err := txn.Commit(); err != nil {
		panic(err)
	}

	fmt.Println("received request to assign delivery agent")
}

func connectToDB(dbName string) (*sql.DB, error) {
	// Replace with your database credentials
	dbUser := "keerthanasmac"
	dbHost := "localhost"
	dbPort := "5432"

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}
