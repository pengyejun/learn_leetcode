package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		root := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if root.Left == nil && root.Right == nil {
			k -= 1
			if k == 0 {
				return root.Val
			}
			continue
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		stack = append(stack, root)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		root.Left = nil
		root.Right = nil
	}
	return 0
}

func main() {

}
