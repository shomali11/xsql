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
	dataSource := fmt.Sprintf(dataSourceFormat, "<USER>", "<PASSWORD>", "<DATABASE_NAME>")
	db, err := sql.Open("<DRIVER>", dataSource)
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
