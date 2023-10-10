package main

import (
	"fmt"
	"optional-patern/db"
)

func main() {
	db := db.New(db.WithHost("192.167.4.3"),
		db.WithPort(5432),
		db.WithDriver("Postgres"),
		db.WithUserNamePassword("admin", "1234"))

	fmt.Println(db)
}
