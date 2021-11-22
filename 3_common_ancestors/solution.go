package main

import (
	"fmt"
)

func main() {
	tree := loadTestData()
	r := solution(tree, 10, 6)
	fmt.Println(r.Value)
}

type TreeNode struct {
	Value    int
	Children []TreeNode
}

func solution(node *TreeNode, nodeA int, nodeB int) TreeNode {
	traceA := &Trace{nodeA, false, TreeNodeStack{data: []TreeNode{}}}
	traceB := &Trace{nodeB, false, TreeNodeStack{data: []TreeNode{}}}
	computeTraces(*node, &Traces{[]*Trace{traceA, traceB}})

	for i, v := range traceA.Path.data {
		if v.Value != traceB.Path.data[i].Value {
			return traceB.Path.data[i-1]
		}
	}

	return TreeNode{}
}

type Trace struct {
	NodeValue    int
	PathComplete bool
	Path         TreeNodeStack
}

func (t *Trace) push(tn TreeNode) {
	t.Path.push(tn)
	if t.NodeValue == tn.Value {
		t.PathComplete = true
	}
}

type Traces struct {
	Data []*Trace
}

func (t *Traces) allCompleted() bool {
	for _, trace := range t.Data {
		if trace.PathComplete == false {
			return false
		}
	}
	return true
}

func (t *Traces) pushPop(pushPop bool, tn TreeNode) {
	for _, trace := range t.Data {
		if !trace.PathComplete {
			if pushPop {
				trace.push(tn)
			} else {
				trace.Path.pop()
			}
		}
	}
}

type TreeNodeStack struct {
	data []TreeNode
}

func (s *TreeNodeStack) pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *TreeNodeStack) push(tn TreeNode) {
	s.data = append(s.data, tn)
}

func computeTraces(root TreeNode, traces *Traces) {
	if traces.allCompleted() {
		return
	}

	traces.pushPop(true, root)
	if traces.allCompleted() {
		return
	}

	if len(root.Children) == 0 {
		traces.pushPop(false, TreeNode{})
	}

	for i, c := range root.Children {
		if traces.allCompleted() {
			return
		}
		computeTraces(c, traces)
		if (len(root.Children)-1) == i && !traces.allCompleted() {
			traces.pushPop(false, TreeNode{})
		}
	}
	return
}

func loadTestData() *TreeNode {
	childrenOfThree := []TreeNode{{7, []TreeNode{}}}
	childrenOfSix := []TreeNode{{8, []TreeNode{}}, {9, []TreeNode{}}}
	childrenOfFive := []TreeNode{{10, []TreeNode{}}}
	childrenOfTwo := []TreeNode{{5, childrenOfFive}, {6, childrenOfSix}}
	children := []TreeNode{{2, childrenOfTwo}, {3, childrenOfThree}, {4, []TreeNode{}}}
	return &TreeNode{1, children}
}
