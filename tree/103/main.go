package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var cur []*TreeNode
	var next []*TreeNode
	var res [][]int
	reversed := false
	cur = append(cur, root)
	for len(cur) > 0 {
		var tmp []int
		for _, v := range cur {
			if reversed {
				tmp = append([]int{v.Val}, tmp...)
			}else {
				tmp = append(tmp, v.Val)
			}
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		reversed = !reversed
		cur = next
		res = append(res, tmp)
		next = []*TreeNode{}
	}
	return res
}

func main() {
	
}
