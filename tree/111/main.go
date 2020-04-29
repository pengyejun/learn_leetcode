package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m1 := minDepth(root.Left)
	m2 := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return m1 + m2 + 1
	}
	if m1 > m2 {
		return m2 + 1
	}else {
		return m1 + 1
	}
}

func main() {
	fmt.Println(minDepth(&TreeNode{Val:1, Left:&TreeNode{Val:2}}))
}
