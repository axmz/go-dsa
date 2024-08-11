package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isDigit(c byte) bool {
	return (c >= '0' && c <= '9')
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z')
}

func buildString(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			sb.WriteByte(s[i])
		} else {
			break
		}

	}

	return sb.String()
}

func buildNumber(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			sb.WriteByte(s[i])
		} else {
			break
		}

	}

	return sb.String()
}

type stack struct {
	data []string
}

func (s *stack) pop() string {
	pop := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return pop
}

func (s *stack) push(i string) {
	s.data = append(s.data, i)
}

func (s *stack) len() int {
	return len(s.data)
}

func (s *stack) shift() string {
	sh := s.data[0]
	s.data = s.data[1:]
	return sh
}

func decodeString(s string) string {
	var sb strings.Builder
	var st stack

	for i := 0; i < len(s); {
		if isDigit(s[i]) {
			n := buildNumber(s[i:])
			st.push(n)
			i += len(n)
		} else if isLetter(s[i]) {
			str := buildString(s[i:])
			st.push(str)
			i += len(str)
		} else if s[i] == '[' {
			st.push("[")
			i++
			continue
		} else if s[i] == ']' {
			newStr := ""
			for p := st.pop(); p != "["; p = st.pop() {
				newStr = p + newStr
			}

			popNum := st.pop()
			n, _ := strconv.Atoi(popNum)
			r := strings.Repeat(newStr, n)
			st.push(r)
			i++
		}

	}

	for st.len() > 0 {
		sb.WriteString(st.shift())
	}

	return sb.String()
}

func main() {
	// s := "3[a]2[bc]"
	// s := "3[a2[c3[x]]]"
	// s := "3[a2[c]]"
	// s := "2[abc]3[cd]ef"
	// s := "abc3[cd]xyz" // "abccdcdcdxyz"
	s := "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	// "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"
	// "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"

	/*

	 */
	fmt.Println(decodeString(s))
}
