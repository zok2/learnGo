package main

import "fmt"

//21. 合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := &ListNode{}
	p := list
	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			p = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return list.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		PrintList(list1)
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list2, list1.Next)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}
func PrintList(list *ListNode) {
	p := list
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func main() {
	List1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	List2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list := mergeTwoLists2(List1, List2)
	PrintList(list)
}
