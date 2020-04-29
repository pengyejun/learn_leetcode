package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func trailver(p, q *TreeNode) bool{
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	return trailver(q.Left, p.Right) && trailver(q.Right, p.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return trailver(root.Left, root.Right)
}

func main() {
	
}
