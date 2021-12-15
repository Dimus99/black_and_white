package main

import (
	"database/sql"
	"fmt"
	"log"
)

func getDBConn() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func writerToDB(inpChan chan [4]string, db *sql.DB, res chan int) {
	for inp := range inpChan {
		pathFrom := inp[0]
		pathTo := inp[1]
		size := inp[2]
		duration := inp[3]
		insertStmt := `insert into "log_entry" values ($1, $2, $3, $4);`
		_, err := db.Exec(insertStmt, pathFrom, pathTo, size, duration)
		if err != nil {
			log.Fatal(err)
		}
	}
	res <- 1
}
