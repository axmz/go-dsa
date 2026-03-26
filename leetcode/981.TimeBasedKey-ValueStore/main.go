package main

import "fmt"

type TimeMap struct {
	data map[string][][2]any
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][][2]any),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], [2]any{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.data[key]; !ok {
		return ""
	}
	values := this.data[key]
	left, right := 0, len(values)-1
	for left <= right {
		mid := left + (right-left)/2
		if values[mid][0].(int) <= timestamp {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 {
		return values[right][1].(string)
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
