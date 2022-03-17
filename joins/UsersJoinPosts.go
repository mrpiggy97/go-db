package joins

import (
	"database/sql"
	"fmt"
)

type UsersInnerJoinPosts struct {
	userId     *int64
	nickname   *string
	postID     *int64
	content    *string
	postUserID *int64
	status     *string
}

func NewUsersInnerJoinPosts() *UsersInnerJoinPosts {
	return &UsersInnerJoinPosts{
		userId:     new(int64),
		nickname:   new(string),
		postID:     new(int64),
		content:    new(string),
		postUserID: new(int64),
		status:     new(string),
	}
}

// UsersJoinPosts makes a join between users table and posts
// and prints the result
func UsersJoinPosts(db *sql.DB) {
	var innerJoinQuery string = "SELECT * FROM users INNER JOIN posts ON users.id=posts.user_id WHERE posts.id > 3 ORDER BY posts.id DESC;"
	data, queryError := db.Query(innerJoinQuery)
	if queryError != nil {
		panic(queryError)
	}
	var queryStructure *UsersInnerJoinPosts = NewUsersInnerJoinPosts()
	for data.Next() {
		var scanningErr error = data.Scan(
			queryStructure.userId,
			queryStructure.nickname,
			queryStructure.postID,
			queryStructure.content,
			queryStructure.postUserID,
			queryStructure.status,
		)

		if scanningErr != nil {
			panic(scanningErr)
		}
		result := fmt.Sprintf(
			"%v,%v,%v,%v,%v,%v",
			*queryStructure.userId,
			*queryStructure.nickname,
			*queryStructure.postID,
			*queryStructure.content,
			*queryStructure.postUserID,
			*queryStructure.status,
		)

		fmt.Println(result)
	}
}
