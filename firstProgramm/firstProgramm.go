package main

import (
    "fmt"
    "os"
)

//var gopath string = os.gopath

func main() {
    fmt.Println(os.Getenv("GOPATH"))
    fmt.Println(os.Getuid())
}