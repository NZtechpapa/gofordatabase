//package mssqlgorm
package mssqlgorm

// Replace with your own connection parameters
var (
    server   = "localhost"
    port     = 14330
    user     = "sa"
    password = "Lei123456"
    database = "barfootDB"
)

func main() {
    // Build connection string
    //  connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
    //     server, user, password, port, database)

    db, err := gorm.Open("mssql", "sqlserver://sa:Lei123456@localhost:14330?database=Master")
    defer db.Close()

    if err != nil {
        panic("failed to connect database")
    }
    db.DropTable(&User{})   // Delete old table if it exists.
    db.CreateTable(&User{}) // Create a new table.

    fmt.Printf("Connected!\n")

    // Add(Create) data
    db.Create(&user)

    fmt.Println(user)
}
