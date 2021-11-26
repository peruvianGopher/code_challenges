package main

import (
    "fmt"
    "sort"
	"time"
)

func main()  {
    fmt.Println(solution([]string{"1+2"}))
}

type Expresion interface {
	result() int
}

type StringMultiplication struct {
	expresion string
}

type StringAddition struct {
	expresion string
}

func (s StringMultiplication) result() int {
	a := Atoi(string(s.expresion[0]))
	b := Atoi(string(s.expresion[2]))
	return a * b
}

func (s StringAddition) result() int {
	a := Atoi(string(s.expresion[0]))
	b := Atoi(string(s.expresion[2]))
	return a + b
}

func getResult(e Expresion) int {
	return e.result()
}

func generateVal(channel chan int, query string) {
	time.Sleep(500 * time.Millisecond)
	val := 0
	if query[1] == '+' {
		val = getResult(StringAddition{query})
	} else {
		val = getResult(StringMultiplication{query})
	}
	channel <- val
}

func solution(queries []string) []int {
	channel := make(chan int)
	for i := 0; i < 1; i++ {
		go generateVal(channel, queries[i])
		queries[i] = ""
	}

	result := []int{}
    for i := 0; i < 1; i++ {
        result = append(result, <-channel)
    }

	sort.Ints(result)
	return result
}

const intSize = 32 << (^uint(0) >> 63)

func Atoi(s string) int {
	const fnAtoi = "Atoi"

	sLen := len(s)
	if intSize == 32 && (0 < sLen && sLen < 10) ||
		intSize == 64 && (0 < sLen && sLen < 19) {
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0
			}
		}

		n := 0
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n
	}

	return 0
}
