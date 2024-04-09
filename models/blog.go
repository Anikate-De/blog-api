package models

import (
	"time"

	"de.anikate/blog-api/db"
)

type Blog struct {
	Id        int64     `json:"id"`
	Title     string    `json:"string" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int64     `json:"likes"`
	Shares    int64     `json:"shares"`
	AuthorId  int64     `json:"author_id" binding:"required"`
}

func AllBlogs() ([]Blog, error) {
	query := `select * from blog;`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blogs []Blog = []Blog{}
	for rows.Next() {
		var blog Blog

		err = rows.Scan(&blog)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, err
}
