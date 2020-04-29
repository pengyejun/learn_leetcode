package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValid1(root *TreeNode) bool{
	if root == nil {
		return true
	}
	if !isValid1(root.Left) {
		return false
	}
	if root.Left != nil && root.Left.Val >= root.Val{
		return false
	}
	if !isValid1(root.Right) {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return true
}

func isValidBST1(root *TreeNode) bool {
	return isValid1(root)
}


func isValid(root *TreeNode, preValue *int) bool{
	if root == nil {
		return true
	}
	if !isValid(root.Left, preValue) {
		return false
	}
	if root.Val <= *preValue {
		return false
	}
	*preValue = root.Val
	if !isValid(root.Right, preValue) {
		return false
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	preValue := math.MinInt64
	return isValid(root, &preValue)
}


func main() {
	t := &TreeNode{Val:10, Left:&TreeNode{Val:5}, Right:&TreeNode{Val:15, Left:&TreeNode{Val:6}, Right:&TreeNode{Val:20}}}
	fmt.Println(isValidBST1(t))
}
