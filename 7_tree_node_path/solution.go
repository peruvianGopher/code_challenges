package main

import "fmt"

func main()  {
    inputA := []int{1,2,4,4,3}
    inputB := []int{2,3,1,3,1}
    solution(4, inputA, inputB) // true

    inputA = []int{1,2,1,3}
    inputB = []int{2,4,3,4}
    solution(4, inputA, inputB) // false

    inputA = []int{2,4,5,3}
    inputB = []int{3,5,6,4}
    solution(6, inputA, inputB) // false

    inputA = []int{1,3}
    inputB = []int{2,2}
    solution(3, inputA, inputB) // true

}

type Node struct {
    Value int
    Pointers map[int]bool
}

func solution(N int, A []int, B []int)  {
    var nodes = map[int]Node{}

    for i, _ := range A {
        if _, ok := nodes[A[i]]; !ok {
            nodes[A[i]] = Node{A[i], map[int]bool{}}
        }

        if _, ok := nodes[B[i]]; !ok {
            nodes[B[i]] = Node{B[i], map[int]bool{}}
        }

        if _, ok := nodes[A[i]].Pointers[B[i]]; !ok {
            nodes[A[i]].Pointers[B[i]] = true
        }

        if _, ok := nodes[B[i]].Pointers[A[i]]; !ok {
            nodes[B[i]].Pointers[A[i]] = true
        }
    }

    fmt.Println(nodes)
    fmt.Println(checkPath(N, nodes))
}

func checkPath(searchNode int, nodes map[int]Node) bool {
   return checkNode(searchNode, nodes[1], nodes, 0)
}

func checkNode(searchNode int, node Node, nodes map[int]Node, ancestor int) bool {
    if node.Value == searchNode {
        return true
    }
    for i, _ := range node.Pointers {
        if (ancestor != nodes[i].Value) && node.Value + 1 == i {
            return checkNode(searchNode, nodes[i], nodes, node.Value)
        }
    }
    return false
}