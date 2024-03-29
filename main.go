package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	groupBy "github.com/mrpiggy97/go-db/groupby"
	insertdata "github.com/mrpiggy97/go-db/insertData"
	"github.com/mrpiggy97/go-db/joins"
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
	insertdata.CreateUser(db)
	joins.UsersJoinPosts(db)
	groupBy.PostsByStatus(db)
	fmt.Println("and that's all fellas")
	fmt.Println("this project used a database created inside a container")
	defer db.Close()
}
