package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update users set name = ? where id = ?")

	stmt.Exec("Breno Machado", 1)
	stmt.Exec("Vanessa Oliveira", 3)

}
