package main

import (
	"database/sql"
	"fmt"
	"github.com/shomali11/xsql"
	"log"
)

const (
	dataSourceFormat = "user=%s password=%s dbname=%s sslmode=disable"
)

func main() {
	dataSource := fmt.Sprintf(dataSourceFormat, "shomali", "", "shomali")
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		log.Fatal(err)
	}

	results, err := xsql.PrettySql(rows)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(results)
}
