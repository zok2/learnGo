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

	candidates1 := []int{2, 3, 6, 7}
	target1 := 7
	r1 := Sum(candidates1, target1)
	fmt.Println(r1)

	candidates2 := []int{2, 3, 5}
	target2 := 8
	r2 := Sum(candidates2, target2)
	fmt.Println(r2)

	candidates3 := []int{2}
	target3 := 1
	r3 := Sum(candidates3, target3)
	fmt.Println(r3)

	//var head = &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//test := reverse(head)
	//for test != nil {
	//	fmt.Println(test.Val)
	//	test = test.Next
	//}
	//l1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 4,
	//		},
	//	},
	//}
	//l2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val: 4,
	//		},
	//	},
	//}
	//List := mergeList(l1, l2)
	//for List != nil {
	//	fmt.Println(List.Val)
	//	List = List.Next
	//}
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
func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
//candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
//对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//
//示例 1：
//
//输入：candidates = [2,3,6,7], target = 7
//输出：[[2,2,3],[7]]
//解释：
//2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//7 也是一个候选， 7 = 7 。
//仅有这两种组合。
//示例 2：
//
//输入: candidates = [2,3,5], target = 8
//输出: [[2,2,2,2],[2,3,3],[3,5]]
//示例 3：
//
//输入: candidates = [2], target = 1
//输出: []
//提示：
//
//1 <= candidates.length <= 30
//2 <= candidates[i] <= 40
//candidates 的所有元素 互不相同
//1 <= target <= 40

func Sum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}
