package main

import "fmt"

var testInput = map[int]int{
    1:1,
}

var testA = []int{1,2}
var testB = []int{1,2}

func main()  {
    for i := range testA {
        fmt.Println(solution(testA[i], testB[i]))
    }
}

func solution(a int, b int) int {
    return a + b
}