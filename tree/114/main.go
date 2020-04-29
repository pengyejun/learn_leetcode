package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func flatten(root *TreeNode)  {
	for root != nil{
		if root.Left != nil{
			right := root.Left
			for right.Right != nil{
				right = right.Right
			}
			right.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}

func main() {
	t := &TreeNode{Val:1, Left:&TreeNode{Val:3, Left:&TreeNode{Val:5}, Right:&TreeNode{Val:6}}, Right:&TreeNode{Val:2}}
	flatten(t)
	fmt.Println(t)
}
