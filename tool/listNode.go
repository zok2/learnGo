package tool

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) int {
	lenght := 0
	for ; head != nil; head = head.Next {
		lenght++
	}
	return lenght
}
func addListNode(head *ListNode, val int) {

}
func PrintList(list *ListNode) {
	p := list
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}
