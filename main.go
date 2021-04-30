package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "siwei.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		_ = db.Close()
	}()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("create table User(Name text, Age integer )")
	result, err := db.Exec("insert into User(`Name`, `Age`) values (?, ?)", "Sam1", 21)
	if err != nil {
		log.Fatalln(err)
	}
	affected, _ := result.RowsAffected()
	log.Println(affected)
	row := db.QueryRow("select * from User limit 1")
	name := ""
	age := 0
	if err := row.Scan(&name, &age); err != nil {
		log.Fatalln(err)
	}
	log.Println(name, age)
}
