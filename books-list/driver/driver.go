package drivers

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mo0zakhraf"
	dbname   = "books"
)

var (
	// Db var
	db *sql.DB
	// DbErr var
	dbErr error
)

// ConnectDB func
func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, dbErr = sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		panic(dbErr)
	}
	//defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return db
}
