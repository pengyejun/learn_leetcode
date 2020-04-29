package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var cur []*TreeNode
	var next []*TreeNode
	var res [][]int
	cur = append(cur, root)
	for len(cur) > 0 {
		var tmp []int
		for _, v := range cur {
			tmp = append(tmp, v.Val)
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		res = append([][]int{tmp}, res...)
		cur = next
		next = []*TreeNode{}
	}
	return res
}


func main() {
	
}
