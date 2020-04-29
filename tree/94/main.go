package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var slice []int

type ListNode struct {
	Val int
	Next *ListNode
}

var l *ListNode
var num int

func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	traversal(root.Left)
	slice = append(slice, root.Val)
	traversal(root.Right)
}

func traversal1(root *TreeNode) {
	if root == nil {
		return
	}
	traversal1(root.Left)
	l.Next = &ListNode{Val:root.Val}
	l = l.Next
	num++
	traversal1(root.Right)
}

func inorderTraversal1(root *TreeNode) []int {
	l = new(ListNode)
	p := l
	num = 0
	traversal1(root)
	p = p.Next
	slice := make([]int, num)
	i := 0
	for p != nil {
		slice[i] = p.Val
		i++
		p = p.Next
	}
	return slice
}

func inorderTraversal(root *TreeNode) []int {
	slice = make([]int, 0)
	traversal(root)
	return slice
}

func main() {
	t := &TreeNode{Val:3, Left:&TreeNode{Val:1}, Right:&TreeNode{Val:4, Left:&TreeNode{Val:2}}}
	fmt.Println(inorderTraversal1(t))
}
