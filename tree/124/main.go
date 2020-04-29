package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var min = math.MinInt64

func maxNodeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxNodeSum(root.Left)
	right := maxNodeSum(root.Right)
	min = max(min, left + right + root.Val)
	return root.Val
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0 ,maxPathSum(root.Left))
	right := max(0, maxPathSum(root.Right))
	min = max(min, left + right + root.Val)
	return max(left, right) + root.Val
}



func main() {
	t := &TreeNode{Val:10, Left:&TreeNode{Val:9, Left:&TreeNode{Val:-7, Left:nil, Right:nil},
		Right:&TreeNode{Val:-2, Left:nil, Right:nil}}, Right:&TreeNode{Val:20,
			Left:&TreeNode{Val:15, Left:nil, Right:nil}, Right:&TreeNode{Val:7, Left:nil, Right:nil}}}
	maxPathSum(t)
	fmt.Println(min)
	min = math.MinInt64
	maxNodeSum(t)
	fmt.Println(min)
}