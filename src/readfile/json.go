//package mssql
package main

import (
    "bufio"
    "fmt"
    "github.com/json-iterator/go"
    "os"
    "strings"
)

func json_read() (fp *os.File) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(dir)
    // Open our jsonFile
    jsonFile, err := os.Open(dir + "testdata.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    fmt.Println("Successfully Opened users.json")
    return jsonFile
}

func json_lineparser(fp *os.File) {
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
