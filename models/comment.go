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

func AllComments(blogId int64) ([]Comment, error) {
	query := `
	select id,
		body, 
		coalesce(parent_id, 0),
		created_at, 
		author_id, 
		blog_id 
	from comment 
	where blog_id = ?;`

	rows, err := db.DB.Query(query, blogId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []Comment = []Comment{}
	for rows.Next() {
		var comment Comment

		err = rows.Scan(
			&comment.Id,
			&comment.Body,
			&comment.ParentId,
			&comment.CreatedAt,
			&comment.AuthorId,
			&comment.BlogId,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, err
}

func GetCommentById(commentId, blogId int64) (*Comment, error) {
	query := `
	select 
		id,
		body,
		coalesce(parent_id, 0),
		created_at,
		author_id,
		blog_id
	from comment 
	where id = ? and blog_id = ?;`

	result := db.DB.QueryRow(query, commentId, blogId)

	var comment Comment

	err := result.Scan(
		&comment.Id,
		&comment.Body,
		&comment.ParentId,
		&comment.CreatedAt,
		&comment.AuthorId,
		&comment.BlogId,
	)
	if err != nil {
		return nil, err
	}

	return &comment, err
}

func (comment *Comment) Update() error {
	query := `
	update comment
	set body = ?
	where id = ?;
	`

	_, err := db.DB.Exec(query, comment.Body, comment.Id)
	return err
}
