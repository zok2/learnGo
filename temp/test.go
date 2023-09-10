package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	L   *TreeNode
	R   *TreeNode
	Val int
}

func main() {
	var head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	test := reverse(head)
	for test != nil {
		fmt.Println(test.Val)
		test = test.Next
	}
}

func test(tNode *TreeNode) {
	if tNode == nil {
		return
	}
	test(tNode.L)
	fmt.Println(tNode.Val)
	test(tNode.R)
}

func reverse(head *ListNode) *ListNode {
	var p *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	return p
}
