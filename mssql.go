//package mssql
package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mssql"
)

// Replace with your own connection parameters
var server = "localhost"
var port = 14330
var user = "SA"
var password = "Lei123456"
var database = "Master"

func main() {
    // Build connection string
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, database)

    var err error

    db, err := gorm.Open("sqlserver", connString)

    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    db.DropTable(&User{})   // Delete old table if it exists.
    db.CreateTable(&User{}) // Create a new table.
    fmt.Printf("Connected!\n")

    user := User{
        Username:  "cihanozhan",
        FirstName: "Cihan",
        LastName:  "Özhan",
    }

    // Add(Create) data
    db.Create(&user)

    fmt.Println(user)
}

type User struct {
    ID        int
    Username  string
    FirstName string
    LastName  string
}
