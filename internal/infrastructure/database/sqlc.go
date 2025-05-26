package database

import (
	"database/sql"

	"github.com/alfariiizi/go-service/database/db"
	_ "github.com/lib/pq"
)

func CreateSQLCDB() *db.Queries {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	queries := db.New(conn)

	return queries
}
