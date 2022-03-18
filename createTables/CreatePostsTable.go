package createtables

import "database/sql"

// CreatePostsTable will create a table called posts
func CreatePostsTable(db *sql.DB) {
	// CreateUsersTable has to run before so tables can be run appropiately
	var query string = "CREATE TABLE posts(id SERIAL PRIMARY KEY, content TEXT NOT NULL, user_id INT NOT NULL, constraint fk_user FOREIGN KEY(user_id) references users(id) ON DELETE CASCADE, status VARCHAR(100) NOT NULL);"
	_, dataErr := db.Query(query)
	if dataErr != nil {
		panic(dataErr)
	}
}
