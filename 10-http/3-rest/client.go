package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserHandler get request and call the right function
func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userByID(w, r, id)
	case r.Method == "GET":
		usersAll(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Sorry...")
	}
}

func userByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:secret@/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var u User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&u.ID, &u.Name)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, string(json))
}

func usersAll(w http.ResponseWriter, r *http.Request) {
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
	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, string(json))
}
