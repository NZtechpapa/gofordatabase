//package mssql
package main
import (
	
     _ "github.com/denisenkom/go-mssqldb"
    "database/sql"
    "context"
    "log"
    "fmt"
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
 //   db, err = sql.Open("sqlserver", connString)
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
    ID        uint
    Username  string
    FirstName string
    LastName  string
}
// CreateEmployee inserts an employee record
func CreateEmployee(name string, location string) (int64, error) {
    ctx := context.Background()
    var err error

    if db == nil {
        err = errors.New("CreateEmployee: db is null")
        return -1, err
    }

    // Check if database is alive.
    err = db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := "INSERT INTO TestSchema.Employees (Name, Location) VALUES (@Name, @Location); select convert(bigint, SCOPE_IDENTITY());"

    stmt, err := db.Prepare(tsql)
    if err != nil {
        return -1, err
    }
    defer stmt.Close()

    row := stmt.QueryRowContext(
        ctx,
        sql.Named("Name", name),
        sql.Named("Location", location))
    var newID int64
    err = row.Scan(&newID)
    if err != nil {
        return -1, err
    }

    return newID, nil
}

// ReadEmployees reads all employee records
func ReadEmployees() (int, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

    // Execute query
    rows, err := db.QueryContext(ctx, tsql)
    if err != nil {
        return -1, err
    }

    defer rows.Close()

    var count int

    // Iterate through the result set.
    for rows.Next() {
        var name, location string
        var id int

        // Get values from row.
        err := rows.Scan(&id, &name, &location)
        if err != nil {
            return -1, err
        }

        fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
        count++
    }

    return count, nil
}

// UpdateEmployee updates an employee's information
func UpdateEmployee(name string, location string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = @Location WHERE Name = @Name")

    // Execute non-query with named parameters
    result, err := db.ExecContext(
        ctx,
        tsql,
        sql.Named("Location", location),
        sql.Named("Name", name))
    if err != nil {
        return -1, err
    }

    return result.RowsAffected()
}

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(name string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name = @Name;")

    // Execute non-query with named parameters
    result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
    if err != nil {
        return -1, err
    }

    return result.RowsAffected()
}