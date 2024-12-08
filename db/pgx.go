package main

import (
    "fmt"
    "log"
    "github.com/jackc/pgx/v4"
    "context"
)

func main() {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
    if err != nil {
        log.Fatal("Unable to connect to database: ", err)
    }
    defer conn.Close(context.Background())

    fmt.Println("Successfully connected to the database!")

    rows, err := conn.Query(context.Background(), "SELECT * FROM products") // replace 'products' with your table name
    if err != nil {
        log.Fatal("Query failed: ", err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        var price float64
        err := rows.Scan(&id, &name, &price) // Replace with your actual column names
        if err != nil {
            log.Fatal("Scan failed: ", err)
        }
        fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", id, name, price)
    }

    if rows.Err() != nil {
        log.Fatal("Row iteration error: ", rows.Err())
    }
}
