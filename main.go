package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : lit <command>")
		os.Exit(1)
	} else {
		fmt.Print(os.Args)
	}
}
