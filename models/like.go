package models

import (
	"time"

	"de.anikate/blog-api/db"
)

type Like struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	AuthorId  int64     `json:"author_id" binding:"required"`
	BlogId    int64     `json:"blog_id" binding:"required"`
}

func (like *Like) Exists() bool {
	query := `
	select id from heart where blog_id = ? and author_id = ?;
	`

	row := db.DB.QueryRow(query, like.BlogId, like.AuthorId)

	return row.Scan(&like.Id) == nil

}

func (like *Like) Save() error {
	query := `
	insert into heart(blog_id, author_id, created_at)
	values (?, ?, datetime('now'));
	update blog set likes = likes + 1 WHERE id = ?;  
	`

	result, err := db.DB.Exec(query, like.BlogId, like.AuthorId, like.BlogId)
	if err != nil {
		return err
	}

	like.Id, err = result.LastInsertId()
	return err
}

func (like *Like) Delete() error {
	query := `
	delete from heart where blog_id = ? and author_id = ?;
	update blog set likes = likes - 1 WHERE id = ?;  
	`

	_, err := db.DB.Exec(query, like.BlogId, like.AuthorId, like.BlogId)
	return err
}

func AllLikes(blogId int64) ([]Like, error) {
	query := `select * from heart where blog_id = ?;`

	rows, err := db.DB.Query(query, blogId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var likes []Like = []Like{}
	for rows.Next() {
		var like Like
		err = rows.Scan(
			&like.Id,
			&like.CreatedAt,
			&like.BlogId,
			&like.AuthorId,
		)
		if err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	return likes, err
}
