package models

import (
	"time"

	"de.anikate/blog-api/db"
)

type Comment struct {
	Id        int64     `json:"id"`
	Body      string    `json:"body" binding:"required"`
	ParentId  int64     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	AuthorId  int64     `json:"author_id" binding:"required"`
	BlogId    int64     `json:"blog_id" binding:"required"`
}

func (comment *Comment) Save() error {
	query := `
	insert into comment(
		body,
		parent_id,
		created_at,
		author_id,
		blog_id
	) values (?, ?, datetime('now'), ?, ?);  
	`

	result, err := db.DB.Exec(query,
		comment.Body,
		func() any {
			if comment.ParentId == 0 {
				return nil
			} else {
				return comment.ParentId
			}
		}(),
		comment.AuthorId,
		comment.BlogId,
	)
	if err != nil {
		return err
	}

	comment.Id, err = result.LastInsertId()
	return err
}
