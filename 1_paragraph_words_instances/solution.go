package main

import (
	"fmt"
	"unicode"
)

const input string = "When!      you're down and low, lower than the floor, And you feel like you ain't got a chance. Bom, bom, bom, Don't make a move till you're in the groove And do the Peter Panda Dance:"

func main() {
	fmt.Println(solution(input))
}

const separator rune = ' '

var notAllowedChars = map[rune]bool{
	',': true,
	';': true,
	'.': true,
	'!': true,
	':': true,
}

func solution(p string) map[string]int {
	counter := make(map[string]int)
	var word []rune
	for _, char := range append([]rune(p), separator) {
		if char != separator {
			if !notAllowedChars[char] {
				word = append(word, unicode.ToLower(char))
			}
			continue
		}

		if len(word) == 0 {
			continue
		}

		counter[string(word)]++
		word = []rune{}
	}
	return counter
}
