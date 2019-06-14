package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var count int

	db, err := sql.Open("postgres", "user=test password=test dbname=foo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(*) FROM table_name")
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}