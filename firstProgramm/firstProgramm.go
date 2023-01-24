package main

import (
	"fmt"
	"os"
	"Github.com/SergeyMoshkov/Formatinput"
)

var gopath string = os.Getenv("GOPATH")

func main() {
	fmt.Println(gopath)
	fmt.Println(os.Getuid())
	FormatInput()
}
