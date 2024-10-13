package main

import (
	"fmt"
	"os"
)

func main () {
	args := os.Args
	for i := 0; i < len(args); i++ {
		fmt.Println("arg[", i, "] = ", args[i])
	}
}