//package mssql
package main

import "github.com/json-iterator/go"
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("testdata.json")
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
        // fmt.Println(arefun)
        fmt.Println("\n")
        var json = jsoniter.ConfigCompatibleWithStandardLibrary
        b, _ := json.Marshal(arefun)
        fmt.Println(b)
    }
    fmt.Println("Successfully Opened users.json")

}
