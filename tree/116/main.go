package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var cur []*Node
	var next []*Node
	cur = append(cur, root)
	for len(cur) > 0 {
		for i, v := range cur {
			if i != len(cur) - 1{
				v.Next = cur[i+1]
			}else {
				v.Next = nil
			}
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		cur = next
		next = []*Node{}
	}
	return root
}

func main() {
	
}
