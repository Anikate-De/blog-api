package models

import (
	"time"

	"de.anikate/blog-api/db"
	"de.anikate/blog-api/utils"
)

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	About     string    `json:"about"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Save() error {

	hash, err := utils.GetHashedPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash

	query := `
	insert into user(name, email, password, about, created_at)
	values(?, ?, ?, ?, datetime('now'));
	`

	result, err := db.DB.Exec(query, u.Name, u.Email, u.Password, u.About)
	if err != nil {
		return err
	}

	u.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Authenticate() error {
	query := `select password from user where email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var hash string
	err := row.Scan(&hash)
	if err != nil {
		return err
	}

	return utils.CompareHash(u.Password, hash)
}
