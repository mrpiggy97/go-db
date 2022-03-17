package groupBy

import (
	"database/sql"
	"fmt"
)

type Post struct {
	status *string
	count  *int64
}

func NewPost() *Post {
	return &Post{
		status: new(string),
		count:  new(int64),
	}
}

// PostsByStatus will get all posts and group them by
// their property status
func PostsByStatus(db *sql.DB) {
	var query string = "SELECT status,count(*) AS posts_count FROM posts GROUP BY status ORDER BY status ASC;"
	data, queryError := db.Query(query)
	if queryError != nil {
		panic(queryError)
	}
	var posts *Post = NewPost()
	for data.Next() {
		var scanningError error = data.Scan(posts.status, posts.count)
		if scanningError != nil {
			panic(scanningError)
		}
		fmt.Println(*posts.status, *posts.count)
	}
}
