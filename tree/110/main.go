package main

import (
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


//递归
func isBalanced(node *TreeNode) bool {
	return find(node) != -1
}

func find(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	l := find(node.Left) //!左节点
	if l == -1 { //剪枝，不平衡时直接返回，
		return -1
	}

	r := find(node.Right)//!右节点
	if r == -1 { //剪枝，不平衡时直接返回
		return -1
	}
	if math.Abs(l-r) > 1 { //剪枝，不平衡时直接返回
		return -1
	}
	return math.Max(l, r) + 1 //计算深度 !根节点
}


func main() {

}
