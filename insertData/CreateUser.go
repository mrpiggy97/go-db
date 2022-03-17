package insertdata

import (
	"database/sql"
	"fmt"
)

func CreateUser(db *sql.DB) {
	var nickname *string = new(string)
	fmt.Println("enter nickname:")
	fmt.Scanln(nickname)
	var query string = fmt.Sprintf("INSERT INTO users(nickname)VALUES('%v');", *nickname)
	data, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.ColumnTypes())
	data, err = db.Query("SELECT * FROM users;")
	if err != nil {
		panic(err)
	}
	for data.Next() {
		var id *int64 = new(int64)
		var nickname *string = new(string)
		var scanningErr error = data.Scan(id, nickname)
		if scanningErr != nil {
			panic(scanningErr)
		}
		fmt.Println(*id, *nickname)
	}
}
