package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

//遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
			continue
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
			continue
		}
		return root
	}
}

func main() {
	
}
