package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	slice := make([]int, 0)
	var cur []*TreeNode
	var next []*TreeNode
	cur = append(cur, root)
	for len(cur) > 0 {
		for i, v := range cur {
			if i == 0 {
				slice = append(slice, v.Val)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
			if v.Left != nil {
				next = append(next, v.Left)
			}
		}
		cur = next
		next = []*TreeNode{}
	}
	return slice
}


func main() {
	
}
