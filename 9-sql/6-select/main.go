package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:secret@/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, errrows := db.Query("select id, name from users")
	if errrows != nil {
		panic(errrows)
	}

	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
	}

}
