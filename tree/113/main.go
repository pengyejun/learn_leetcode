package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var slice [][]int

func depth(root *TreeNode, sum int, res []int) {
	if root == nil {
		return
	}
	sum -= root.Val
	res = append(res, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			tmpRes := make([]int, len(res))
			copy(tmpRes, res)
			slice = append(slice, tmpRes)
		}
		return
	}
	depth(root.Left, sum, res)
	depth(root.Right, sum, res)
}

func pathSum(root *TreeNode, sum int) [][]int {
	slice = make([][]int, 0)
	depth(root, sum, []int{})
	return slice
}

func abc(res []int, i int) {
	if i < 2 {
		i += 1
		tmp_res := make([]int, len(res))
		copy(tmp_res, res)
		tmp_res = append(tmp_res, 2)
		abc(tmp_res, i)
		return
	}
	fmt.Println(res)
}

func main() {
	abc([]int{}, 0)
}
