package main

import (
	"log"

	"github.com/gustavokuklinski/myconn"
)

func main() {
	db := myconn.DbConn("YOUR USERNAME", "YOUR PASSWORD", "YOUR DATABASE")

	_, err := db.Query("YOUR SQL QUERY")

	if err != nil {
		panic(err.Error())
	}

	log.Println("Succefull conected")

	defer db.Close()
}
