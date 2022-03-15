package main

import (
	"fmt"

)

type Tree struct {
	data       string
	leftChild  *Tree
	rightChild * Tree
}
//CreateTree 递归创建
func CreateTree () *Tree {
		var value string
		fmt.Scanln(&value)
		node := new(Tree)
		if value != "0" {
			node.data = value
			node.leftChild = CreateTree()
			node.rightChild = CreateTree()
		}else {
			return nil
		}
		return node
}

func inorderTraversal(root *Tree) (res []string) {
	var inorder func(node *Tree)
	inorder = func(node *Tree) {
		if node == nil {
			return
		}
		inorder(node.leftChild)
		res = append(res, node.data)
		inorder(node.rightChild)
	}
	inorder(root)
	return
}


func main()  {
	root := CreateTree()
	data := inorderTraversal(root)
	fmt.Println(data)
}