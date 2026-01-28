package main

import "fmt"

type Logger struct {
	m map[string]int
}

func Constructor() Logger {
	return Logger{m: make(map[string]int)}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if t, ok := this.m[message]; ok {
		if timestamp-t < 10 {
			return false
		}
	}
	this.m[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
