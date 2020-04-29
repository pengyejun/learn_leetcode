package main

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var slice []string

func depth(root *TreeNode, path string) {
	if root == nil {
		return
	}
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		slice = append(slice, path)
	}else{
		path += "->"
		depth(root.Left, path)
		depth(root.Right, path)
	}
}


func binaryTreePaths(root *TreeNode) []string {
	slice =  make([]string, 0)
	depth(root, "")
	return slice
}

func main() {
}
