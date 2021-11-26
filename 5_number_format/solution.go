package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "1234567890123"
	//input := "0 - 22 1985--324"
    //input := "555372654"
	fmt.Println(solutionb(input))
}

func solutionb(n string) string {
	runes := []rune(n)
	var response []rune
	for i := 0; i < len(runes); i++ {
		if !unicode.IsNumber(runes[i]) {
			runes = append(runes[:i], runes[i+1:]...)
			i = i - 1
		}
	}

	for i := 0; i < len(runes); i++ {
		if (i+1)%3 == 0 && i < len(runes) - 1 {
			if (len(runes) - (i + 1)) >= 2 {
				response = append(response, runes[i])
				response = append(response, '-')
			} else {
				response = append(response, '-')
				response = append(response, runes[i])
			}
		} else {
			response = append(response, runes[i])
		}
	}
	return string(response)
}