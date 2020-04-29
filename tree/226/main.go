package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func swap(root *TreeNode) {
	if root == nil{
		return
	}
	root.Left, root.Right = root.Right, root.Left
	swap(root.Left)
	swap(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}
	swap(root)
	return root
}

func main() {
	t := invertTree(&TreeNode{Val:1, Left:&TreeNode{Val:3, Left:&TreeNode{Val:5}, Right:&TreeNode{Val:6}}, Right:&TreeNode{Val:2}})
	fmt.Println(t)
}
