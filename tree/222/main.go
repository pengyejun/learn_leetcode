package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func depth(root *TreeNode) {
	if root == nil {
		return
	}
	res += 1
	depth(root.Left)
	depth(root.Right)
}

func countNodes(root *TreeNode) int {
	res = 0
	depth(root)
	return res
}

func main() {
	
}
