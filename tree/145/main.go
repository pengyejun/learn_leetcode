package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	slice := make([]int, 0)
	for len(stack) > 0 {
		root := stack[len(stack) - 1]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Left == nil && root.Right == nil {
			stack = stack[:len(stack) - 1]
			slice = append(slice, root.Val)
		}
		root.Left = nil
		root.Right = nil
	}
	return slice
}

func main() {
	
}
