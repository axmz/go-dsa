package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	t := time.Now()
	go func() {
		tt := time.Now().Sub(t)
		fmt.Println("end go", tt)
	}()

	for {

	}
	fmt.Println("end")
}
