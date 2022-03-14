package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func main() {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var query string = "select * from posts where id > 4;"
	data, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var content *string = new(string)
	var id *int64 = new(int64)
	var status *string = new(string)
	var user_id *string = new(string)
	for data.Next() {
		var scanningErr error = data.Scan(id, content, user_id, status)
		if scanningErr != nil {
			panic(scanningErr)
		}
		fmt.Println(*id, *content, *status, *user_id)
	}

	var secondQuery string = "select id, count(*) as posts_count from posts where id > 3 group by id order by id desc;"
	data, err = db.Query(secondQuery)
	if err != nil {
		panic(err)
	}
	var postsCount *int = new(int)
	for data.Next() {
		var scanningErr error = data.Scan(id, postsCount)
		if scanningErr != nil {
			panic(scanningErr)
		}
		fmt.Println(*id, *postsCount)
	}
}
