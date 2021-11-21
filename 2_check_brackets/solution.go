package main

import "fmt"

type stack struct {
	data []rune
}

func (s *stack) pop() rune {
	output := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return output
}

func (s *stack) push(str rune) {
	s.data = append(s.data, str)
}

var closePair = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

var inputs = []string{
	"[()]{}{[()()]()}",
	"[(])",
}

func main() {
	for _, i := range inputs {
		fmt.Println(solution(i))
	}
}

func solution(s string) bool {
	stk := stack{data: []rune{}}
	for _, c := range []rune(s) {
		if _, ok := closePair[c]; !ok {
			stk.push(c)
		} else {
			if stk.pop() != closePair[c] {
				return false
			}
		}
	}
	return len(stk.data) == 0
}