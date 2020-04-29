package main

import (
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var preNode, firstNode, secondNode *TreeNode

func recovertree(root *TreeNode) {
	if root == nil {
		return
	}
	recovertree(root.Left)
	if firstNode == nil && preNode.Val >= root.Val {
		firstNode = preNode
	}
	if firstNode != nil && preNode.Val >= root.Val {
		secondNode = root
	}
	preNode = root
	recovertree(root.Right)
}

func recoverTree(root *TreeNode)  {
	firstNode = nil
	preNode = &TreeNode{Val:math.MinInt64}
	recovertree(root)
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}


func main() {
	t := &TreeNode{Val:3, Left:&TreeNode{Val:1}, Right:&TreeNode{Val:4, Left:&TreeNode{Val:2}}}
	recoverTree(t)
}
