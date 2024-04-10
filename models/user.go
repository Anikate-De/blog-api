package models

import (
	"errors"
	"time"

	"de.anikate/blog-api/db"
	"de.anikate/blog-api/utils"
)

type User struct {
	Uid       int64     `json:"uid"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password,omitempty" binding:"required"`
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

	u.Uid, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Authenticate() error {
	query := `select password, uid from user where email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var hash string
	err := row.Scan(&hash, &u.Uid)
	if err != nil {
		return err
	}

	return utils.CompareHash(u.Password, hash)
}

func (user *User) Delete() error {
	query := `
	delete from user
	where uid = ?;
	`

	res, err := db.DB.Exec(query, user.Uid)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}

func GetUserByID(uid int64) (*User, error) {
	query := `select uid, name, email, about, created_at from user where uid = ?;`

	result := db.DB.QueryRow(query, uid)

	var user User

	err := result.Scan(
		&user.Uid,
		&user.Name,
		&user.Email,
		&user.About,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (user *User) Update() error {
	query := `
	update user
	set name = ?, about = ?
	where uid = ?;
	`

	_, err := db.DB.Exec(query, user.Name, user.About, user.Uid)
	return err
}
