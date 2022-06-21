package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into users(name) values (?)")
	stmt.Exec("Breno")
	stmt.Exec("Bruno")

	res, _ := stmt.Exec("Vanessa")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	affectedRows, _ := res.RowsAffected()
	fmt.Println(affectedRows)
}
