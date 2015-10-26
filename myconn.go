// Package myconn connect to MySQL database on a easy beginner way.
package myconn

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Func DbConn get the user, pass and dbname, returning a db object
func DbConn(user, pass, dbname string) (db *sql.DB) {

	// Open the connection
	db, err := sql.Open("mysql", user+":"+pass+"@/"+dbname)

	// If fail, display error
	if err != nil {
		panic(err.Error())
	}

	// Return the *sql.DB object
	return db
}
