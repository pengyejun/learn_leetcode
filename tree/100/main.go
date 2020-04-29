package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}



func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil{
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {

	t := &TreeNode{Val:3, Left:&TreeNode{Val:1}, Right:&TreeNode{Val:4, Left:&TreeNode{Val:2}}}
	t1 := &TreeNode{Val:3, Left:&TreeNode{Val:1}, Right:&TreeNode{Val:4, Left:&TreeNode{Val:2}}}
	t2 := &TreeNode{Val:2}
	t3 := &TreeNode{Val:2}
	t3.Right = &TreeNode{Val:3}
	fmt.Println(t3 == t2)
	fmt.Println(isSameTree(t, t1))
}
