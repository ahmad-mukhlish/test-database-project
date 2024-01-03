package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func main() {

	//connect to the database
	db, err := sql.Open("pgx", "host=localhost port=54321 dbname=test user=ahmadmukhlis password=password")
	if err != nil {
		log.Fatal("error", err)
	}
	defer db.Close()

	//ping the database
	err = db.Ping()
	if err != nil {
		log.Fatal("error", err)
	}

	//show the database content

	err = getRows(db)
	if err != nil {
		log.Fatal("error", err)
	}

	//insert a row to a database

	//show the database content

	//update a row in the database

	//show the database content

	//delete a row in the database

	//show the database content

}

func getRows(db *sql.DB) error {

	selectQuery := `select id, first_name, last_name from users`

	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal("error", err)
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName)
		fmt.Printf("Row %d, %s %s \n", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("error", err)
	}

	fmt.Println("-----------------")

	return nil
}
