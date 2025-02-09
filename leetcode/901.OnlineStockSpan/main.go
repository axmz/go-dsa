package main

import "fmt"

type StockSpanner struct {
	stocks [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stocks: [][2]int{{1, 0}},
	}
}

func (this *StockSpanner) Next(price int) int {
	l := len(this.stocks)
	span := 0
	for i := l - 1; i >= 0; {
		if this.stocks[i][1] <= price {
			span += this.stocks[i][0]
			i -= this.stocks[i][0]
		} else {
			span++
			break
		}
	}
	this.stocks = append(this.stocks, [2]int{span, price})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func main() {

	s := Constructor()
	s.Next(100)
	s.Next(80)
	s.Next(60)
	s.Next(70)
	s.Next(60)
	s.Next(75)
	s.Next(85)
	fmt.Println(s.stocks)
}
