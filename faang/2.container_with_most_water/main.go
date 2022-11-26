package main

import (
	"fmt"
	"os"
)

const (
	http = "http://"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 1 {
		fmt.Println("Give me smth to maks")
		return
	}

	fmt.Println(args)
}
