package main

import (
	"./internal"
)

/*
*

		type ListNode struct {
			Val  int
			Next *ListNode
		}
	 19. 删除链表的倒数第 N 个结点
	    输入：head = [1,2,3,4,5], n = 2
	    输出：[1,2,3,5]
	    - Definition for singly-linked list.
	    - type ListNode struct {
	    - Val int
	    - Next *ListNode
	    - }
*/
func removeNthFromEnd(head *internal.ListNode, n int) *internal.ListNode {
	var d = &internal.ListNode{0, head}
	var l, r = head, d
	for i := 0; i < n; i++ {
		l = l.Next
	}
	for ; l != nil; l = l.Next {
		r = r.Next
	}

	r.Next = r.Next.Next
	return d.Next
}

func removeNtfFromEnd2(head *internal.ListNode, n int) *internal.ListNode {
	var stack []*internal.ListNode
	d := &internal.ListNode{0, head}
	for node := d; node != nil; node = node.Next {
		stack = append(stack, node)
	}
	tmp := stack[len(stack)-1-n]
	tmp.Next = tmp.Next.Next
	return d.Next
}

func main() {

	var list = &internal.ListNode{
		Val: 1,
	}
	var head = list
	addList := func(i int) {
		var node = new(internal.ListNode)
		node.Val = i
		list.Next = node
		list = list.Next
	}
	for i := 2; i <= 5; i++ {
		addList(i)
	}
	//d1 := removeNthFromEnd(head, 3)
	d2 := removeNtfFromEnd2(head, 2)
	//PrintList(d1)
	d2.PrintList(d2)
}
