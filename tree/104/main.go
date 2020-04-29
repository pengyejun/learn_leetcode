package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var max int

func depth(root *TreeNode, res int) {
	if root == nil {
		return
	}
	res += 1
	if res > max {
		max = res
	}
	depth(root.Left, res)
	depth(root.Right, res)
}

func maxDepth(root *TreeNode) int {
	max = 0
	if root == nil {
		return 0
	}
	depth(root, 0)
	return max
}



func main() {
	
}
