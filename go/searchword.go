package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Input string, a paragraph
// Number of instances of each word from the text in the text

const input string = "When!      you're down and low, lower than the floor, And you feel like you ain't got a chance. Bom, bom, bom, Don't make a move till you're in the groove And do the Peter Panda Dance:"

func main() {
	fmt.Println(solution(input))
}

func solution(p string) map[string]int {
	counter := make(map[string]int)
	var word string
	for _, char := range []rune(p+" ") {
		if char != ' ' {
			word = word + string(char)
			continue
		}

		if word == "" {
			continue
		}

		nWord := normalizeWord(word)
		word = ""
		counter[nWord]++
	}
	return counter
}

var normalizedWordsMap = make(map[string]string)
func normalizeWord(word string) string {
	if value, ok := normalizedWordsMap[word]; ok {
		return value
	}
	re := regexp.MustCompile(`[:,!.]`)
	normalizedWordsMap[word] = re.ReplaceAllString(strings.ToLower(word), "")
	return normalizedWordsMap[word]
}