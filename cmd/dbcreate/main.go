package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Create("./db/data.db")
	os.MkdirAll("./db", 0755)

	db, err := sql.Open("sqlite3", "./db/data.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec("CREATE TABLE `messages` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `name` VARCHAR(255) NOT NULL, `message` VARCHAR(255) NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
