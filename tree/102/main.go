package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var cur []*TreeNode
	var next []*TreeNode
	var res [][]int
	cur = append(cur, root)
	for len(cur) > 0 {
		var tmp []int
		for i := 0; i < len(cur); i++ {
			tmp = append(tmp, cur[i].Val)
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		res = append(res, tmp)
		cur = next
		next = []*TreeNode{}
	}
	return res
}

func main() {
	
}
