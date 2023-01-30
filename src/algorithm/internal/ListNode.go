package internal

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (*ListNode) PrintList(list *ListNode) {
	p := list
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}
