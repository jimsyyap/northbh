package pkg

import (
    "fmt"
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

func InitDB(dataSourceName string) {
    var err error
    DB, err = sqlx.Connect("postgres", dataSourceName)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    fmt.Println("Connected to database")
}
