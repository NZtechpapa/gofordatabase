package main

//  import the mssql and xorm
import (
    "./orm/xorm"
    "./orm/xorm/models"
    "bufio"
    "fmt"
    // "github.com/go-xorm/xorm"
    "github.com/json-iterator/go"
    "os"
    "strings"
)

func main() {

    // init the dataset
    xormcode.Xorm_init()
    // Open our jsonFile
    jsonFile, err := os.Open("../testdataset/testdata1.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

    fmt.Println("Successfully Opened users.json")

    scanner := bufio.NewScanner(jsonFile)
    for scanner.Scan() {
        split1 := scanner.Text()
        split := strings.Index(split1, "{")
        arefun := split1[split : len(scanner.Text())-1]

        //   fmt.Println(arefun)
        arfun1 := "[" + arefun + "]"
        //fmt.Println(arfun1)
        //  fmt.Println("\n")
        var json = jsoniter.ConfigCompatibleWithStandardLibrary
        // fmt.Println(json.Valid([]byte(arefun)))
        //  fmt.Println("\n")
        var table1 []models.Barfoottable1
        err := json.Unmarshal([]byte(arfun1), &table1)
        // err := json.Unmarshal(arefun, &table1)
        if err != nil {
            fmt.Println("error:", err)
        }
        //  fmt.Printf("%+v", table1)
        fmt.Printf("%+v", table1[0].DVRValuationImprovement)
        fmt.Println("\n")
        xormcode.Xorm_insert(table1[0])
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

}
