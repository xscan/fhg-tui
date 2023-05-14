package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // nolint: golint
)

// DB holds the database information
type DB struct {
	db *sql.DB
	// c  *Controller
}

// Init setups the database and creates tables if needed.
func (d *DB) Init(dbFile string) error {
	// d.c = c
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println(err)
	}
	d.db = db
	// defer d.db.Close()

	_, err = d.db.Exec(`
         create table if not exists projects(
			id integer not null primary key,
			number text,
			name text,
			desc text,
			category text,
			uri text,
			star text,
			fork text,
			synced DATETIME
		);`)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 关闭数据库
func (d *DB) Close() {
	d.db.Close()
}

// CleanupDB removes old and deleted articles
func (d *DB) CleanupDB() {
	st, err := d.db.Prepare(fmt.Sprintf(
		"delete from projects where 1 =1 "),
	)
	if err != nil {
		log.Println(err)
	}
	defer st.Close()

	if _, err = st.Exec(); err != nil {
		log.Println(err)
	}
}

// All fetches all articles from the database
func (d *DB) All() []ProjectItem {
	st, err := d.db.Prepare("select id,name,number,category,desc,uri,star,fork from projects order by id")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var (
		id       int
		name     string
		category string
		number   string
		uri      string
		desc     string
		star     string
		fork     string
	)

	items := []ProjectItem{}

	for rows.Next() {
		err = rows.Scan(&id, &number, &name, &category, &fork, &uri, &desc, &star)
		if err != nil {
			log.Println(err)
		}

		items = append(items, ProjectItem{
			number:   number,
			uri:      uri,
			desc:     desc,
			category: category,
			star:     star,
			fork:     fork,
			name:     name,
		})
	}
	return items
}

// 获取当前期数
func (d *DB) GetNumber(currentNumber string) []ProjectItem {
	st, err := d.db.Prepare("select id,name,number,category,desc,uri,star,fork from projects where number=? order by id")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer st.Close()

	rows, err := st.Query(currentNumber)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var (
		id       int
		name     string
		category string
		number   string
		uri      string
		desc     string
		star     string
		fork     string
	)

	items := []ProjectItem{}

	for rows.Next() {
		err = rows.Scan(&id, &number, &name, &category, &fork, &uri, &desc, &star)
		if err != nil {
			log.Println(err)
		}

		items = append(items, ProjectItem{
			number:   number,
			uri:      uri,
			desc:     desc,
			category: category,
			star:     star,
			fork:     fork,
			name:     name,
		})
	}
	return items
}

func (d *DB) Save(a ProjectItem) error {
	st, err := d.db.Prepare("select name from projects where name=? order by id")
	if err != nil {
		log.Println(err)
	}
	defer st.Close()

	res, err := st.Query(a.name)
	if err != nil {
		log.Println(err)
	}
	defer res.Close()
	for res.Next() {
		return nil
	}

	tx, err := d.db.Begin()
	if err != nil {
		log.Println(err)
	}

	st, err = tx.Prepare("insert into projects(name, number, category, desc, uri, star, fork,synced) values(?, ?, ?, ?, ?, ?, ?,?)")
	if err != nil {
		log.Println(err)
	}
	defer st.Close()

	if _, err := st.Exec(a.name, a.number, a.category, a.desc, a.uri, a.star, a.fork, time.Now().Format("2006-01-02 15:04:05")); err != nil {
		log.Println(err)
	}

	tx.Commit()
	return nil
}

// Delete marks an article as deleted. Will not remove it from DB (see CleanupDB)
// func (d *DB) Delete(a *ProjectItem) {
// 	st, err := d.db.Prepare("update articles set deleted = true where id = ?")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer st.Close()

// 	if _, err := st.Exec(a.id); err != nil {
// 		log.Println(err)
// 	}
// }

// // MarkRead marks an article as read in the database
// func (d *DB) MarkRead(a *Article) error {
// 	st, err := d.db.Prepare("update articles set read = true where id = ?")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer st.Close()

// 	if _, err := st.Exec(a.id); err != nil {
// 		log.Println(err)
// 	}
// 	return nil
// }

// // MarkUnread marks an article as unread in the database
// func (d *DB) MarkUnread(a Article) error {
// 	st, err := d.db.Prepare("update articles set read = false where id = '?'")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer st.Close()

// 	if err := st.QueryRow(a.id); err != nil {
// 		log.Println(err)
// 	}
// 	return nil
// }

// // MarkAllRead marks all articles in the database as read
// func (d *DB) MarkAllRead(feed string) {
// 	stmt := "update articles set read = true"
// 	if feed != "" {
// 		stmt = "update articles set read = true where feed = ?"
// 	}

// 	st, err := d.db.Prepare(stmt)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer st.Close()

// 	if feed != "" {
// 		if _, err := st.Exec(feed); err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		if _, err := st.Exec(); err != nil {
// 			log.Println(err)
// 		}
// 	}
// }

// // MarkAllUnread marks all articles in the database as not read
// func (d *DB) MarkAllUnread(feed string) {
// 	stmt := "update articles set read = false"
// 	if feed != "" {
// 		stmt = "update articles set read = false where feed = ?"
// 	}

// 	st, err := d.db.Prepare(stmt)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer st.Close()

// 	if feed != "" {
// 		if _, err := st.Exec(feed); err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		if _, err := st.Exec(); err != nil {
// 			log.Println(err)
// 		}
// 	}
// }
