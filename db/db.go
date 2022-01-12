package db_cennetion

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	// "log"
	// "net/http"

	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// const (
//   host     = "localhost"
//   port     = 5438
//   user     = "postgres"
//   password = "postgres"
//   dbname   = "postgres"
// )

// func db_connection() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 	  "password=%s dbname=%s sslmode=disable",
// 	  host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 	  panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 	  panic(err)
// 	}

// 	fmt.Println("Successfully connected!")
// 	return db
//   }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5438
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func sayHello() {
	fmt.Println("Helloe")
}
