package createtables

import "database/sql"

// CreateUsersTable will create users table for database
// if it does not exist
func CreateUsersTable(db *sql.DB) {
	var query string = "CREATE TABLE users(id SERIAL PRIMARY KEY, nickname VARCHAR(255) NOT NULL);"
	_, dataErr := db.Query(query)
	if dataErr != nil {
		panic(dataErr)
	}
}
