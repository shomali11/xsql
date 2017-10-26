# xsql [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/xsql)](https://goreportcard.com/report/github.com/shomali11/xsql) [![GoDoc](https://godoc.org/github.com/shomali11/xsql?status.svg)](https://godoc.org/github.com/shomali11/xsql) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

SQL Query Results Pretty Printing

## Usage

Using `govendor` [github.com/kardianos/govendor](https://github.com/kardianos/govendor):

```
govendor fetch github.com/shomali11/xsql
```

## Dependencies

* `util` [github.com/shomali11/util](https://github.com/shomali11/util)

# Examples

## Example 1

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
```

```
 id  |      name      |         title         |         created_at          | number | decimal | active
-----+----------------+-----------------------+-----------------------------+--------+---------+--------
   1 | Raed Shomali   | Sr. Software Engineer | 2017-10-24T20:59:43.37154Z  |     11 | 789.123 | true
   2 | Dwayne Johnson | The Rock              | 2017-10-24T21:00:31.530534Z |   1000 |     3.7 | true
 300 | Steve Austin   | Stone Cold            | 2017-10-26T19:42:51.993465Z |  55000 |   55.55 | false
(3 rows)
```