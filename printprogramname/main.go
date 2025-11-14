package main

import (
	"fmt"
	"os"
)

func main() {
	progname := os.Args
	fmt.Println(progname[0])
}
