package models

import (
	"errors"
	"time"

	"de.anikate/blog-api/db"
)

type Blog struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title" binding:"required"`
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

		err = rows.Scan(&blog.Id,
			&blog.Title,
			&blog.Content,
			&blog.CreatedAt,
			&blog.Likes,
			&blog.Shares,
			&blog.AuthorId,
		)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, err
}

func GetBlogById(id int64) (*Blog, error) {
	query := `select * from blog where id = ?;`

	result := db.DB.QueryRow(query, id)

	var blog Blog

	err := result.Scan(&blog.Id,
		&blog.Title,
		&blog.Content,
		&blog.CreatedAt,
		&blog.Likes,
		&blog.Shares,
		&blog.AuthorId,
	)
	if err != nil {
		return nil, err
	}

	return &blog, err
}

func (blog *Blog) Save() error {
	query := `
	insert into blog(
		title,
		content,
		created_at,
		author_id
	) values (
		?, ?, datetime('now'), ?
	);  
	`

	result, err := db.DB.Exec(query, blog.Title, blog.Content, blog.AuthorId)
	if err != nil {
		return err
	}

	blog.Id, err = result.LastInsertId()
	return err
}

func (blog *Blog) Update() error {
	query := `
	update blog
	set title = ?, content = ?
	where id = ?;
	`

	_, err := db.DB.Exec(query, blog.Title, blog.Content, blog.AuthorId)
	return err
}

func (blog *Blog) Delete() error {
	query := `
	delete from blog
	where id = ? and author_id = ?;
	`

	res, err := db.DB.Exec(query, blog.Id, blog.AuthorId)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}

func ShareBlog(blogId int64) error {
	query := `
	update blog
	set shares = shares + 1
	where id = ?;
	`

	rows, err := db.DB.Exec(query, blogId)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}
