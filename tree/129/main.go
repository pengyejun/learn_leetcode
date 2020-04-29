package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res int

func depth(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = root.Val + sum * 10
	if root.Left == nil && root.Right == nil {
		res += sum
		return
	}
	depth(root.Left, sum)
	depth(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	res = 0
	if root == nil {
		return 0
	}
	depth(root, 0)
	return res
}

func main() {
	
}
