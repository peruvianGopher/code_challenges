package main

import "fmt"

func main()  {
    input := []int{1,4,0,7,0,0,8,9,0,0,4}
    fmt.Println(solution(input))
}

func solution(list []int) []int {
    swapCounter := 0
    for i, v := range list {
        if v == 0 {
            if swapCounter == 0 {
                swapCounter = i + 1
            }

            for {
                if swapCounter == len(list) {
                    break
                }

                if list[swapCounter] != 0 {
                    list[i] = list[swapCounter]
                    list[swapCounter] = v
                    break
                }
                swapCounter = swapCounter + 1
            }
        }
    }
    return list
}