package main

import (
    "fmt"
    "sort"
	"strings"
)

var testsA = [][]string{
	{"pim", "pom"},
	{"sander", "amy", "ann", "michael"},
	{"adam", "eva", "leo"},
}

var testsB = [][]string{
	{"999999999", "777888999"},
	{"123456789", "234567890", "789123456", "123123123"},
	{"121212121", "111111111", "444555666"},
}

var testsP = []string{"88999", "1", "112"}

func main() {
    for i := range testsA {
        s := solution(testsA[i], testsB[i], testsP[i])
        fmt.Println(s)
    }
}

func solution(A []string, B []string, P string) string {
	var m = map[string]string{}
	for i, s :=  range A {
		m[s] = 	B[i]
	}

	sort.Strings(A)
	for _, s :=  range A {
		if strings.Contains(m[s], P) {
			return s
		}
	}
	return "NO CONTACT"
}