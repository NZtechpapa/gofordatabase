package main

//  import the mssql and xorm
import (
    "bufio"
    "fmt"
    "github.com/go-xorm/xorm"
    "github.com/json-iterator/go"
    "os"
    "strings"
)

func main() {

    // Open our jsonFile
    jsonFile, err := os.Open("../testdataset/testdata.json")
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
        fmt.Println(arefun)
        fmt.Println("\n")

    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    var json = jsoniter.ConfigCompatibleWithStandardLibrary

    json.Marshal(&jsonFile)

}
