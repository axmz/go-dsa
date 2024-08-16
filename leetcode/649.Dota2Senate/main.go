package main

import "fmt"

func predictPartyVictory(senate string) string {
	var Rc, Dc int
	queue := make([]byte, len(senate))

	for i, b := range []byte(senate) {
		queue[i] = b
		if b == 'R' {
			Rc++
		} else {
			Dc++
		}
	}

	var Rv, Dv int
	var cur byte
	for i := 0; Rc > 0 && Dc > 0; i++ {
		cur = queue[i]
		if cur == 'R' {
			if Dv > 0 {
				Rc--
				Dv--
			} else {
				queue = append(queue, 'R')
				Rv++
			}
		} else {
			if Rv > 0 {
				Dc--
				Rv--
			} else {
				queue = append(queue, 'D')
				Dv++
			}
		}
	}

	if Rc == 0 {
		return "Dire"
	}

	return "Radiant"
}

func main() {
	s := "RDDDDRRRR"
	res := predictPartyVictory(s)
	fmt.Println(res)
}
