package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	l := lowestCommonAncestor1(root.Left, p, q)
	r := lowestCommonAncestor1(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}

//遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		root = stack[len(stack) - 1]
		if root.Val == p.Val || root.Val == q.Val {
			
		}
	}
}


func main() {
	
}
