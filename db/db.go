package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect to the SQLite DB
func Connect() {
	if _, err := os.Stat("./out"); os.IsNotExist(err) {
		os.Mkdir("out", os.ModeDir)
	}

	var err error
	DB, err = sql.Open("sqlite3", "out/blog.db")

	if err != nil {
		panic(err)
	}

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)

	log.Printf("Connected to the Database")

	setup()
	createTables()
}

// Sets SQLite Database to use foreign keys for tables
func setup() {
	query := `pragma foreign_keys = ON;`

	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
		panic("unable to setup foreign keys")
	}
}

func createTables() {
	createUserTable()
	createBlogTable()
	createCommentTable()
	createHeartTable()
}

func createUserTable() {
	query := `
	create table if not exists user(
		uid integer primary key autoincrement,
		name text not null,
		email text not null unique,
		password text not null,
		about text,
		created_at datetime not null
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
		panic("unable to create user table")
	}
}

func createBlogTable() {
	query := `
	create table if not exists blog(
		id integer primary key autoincrement,
		title text not null,
		content text not null,
		created_at datetime not null,
		likes integer default 0,
		shares integer default 0,
		author_id integer not null,
		foreign key (author_id) references user(uid) on delete cascade
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
		panic("unable to create blog table")
	}
}

func createCommentTable() {
	query := `
	create table if not exists comment(
		id integer primary key autoincrement,
		body text not null,
		parent_id integer,
		created_at datetime not null,
		author_id integer not null,
		blog_id integer NOT null,
		foreign key (parent_id) references comment(id) on delete cascade,
		foreign key (author_id) references user(uid) on delete cascade,
		foreign key (blog_id) references blog(id) on delete cascade
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
		panic("unable to create comment table")
	}
}

func createHeartTable() {
	query := `
	create table if not exists heart(
		id integer primary key autoincrement,
		created_at datetime NOT null,
		blog_id integer not null,
		author_id integer not null,
		
		foreign key (blog_id) references blog(id) on delete cascade,
		foreign key (author_id) references user(uid) on delete cascade
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
		panic("unable to create heart table")
	}
}
