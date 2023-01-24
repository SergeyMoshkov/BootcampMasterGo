package main

import (
	"fmt"
	"os"
	"github.com/SergeyMoshkov/bootcampMasterGo"
)

var gopath string = os.Getenv("GOPATH")

func main() {
	fmt.Println(gopath)
	fmt.Println(os.Getuid())
	FormatInput()
}
