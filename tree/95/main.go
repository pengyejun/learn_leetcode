package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generatetrees(start, end int) []*TreeNode {
	allTrees := make([]*TreeNode, 0)
	if start > end {
		allTrees = append(allTrees, nil)
		return allTrees
	}

	for i := start; i <= end; i++ {
		leftTrees := generatetrees(start, i - 1)
		rightTrees := generatetrees(i + 1, end)

		for _, l := range leftTrees {
			for _, r := range rightTrees {
				cur := &TreeNode{Val:i}
				cur.Left = l
				cur.Right = r
				allTrees = append(allTrees, cur)
			}
		}
	}
	return allTrees
}


func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generatetrees(1, n)
}

func main() {
	fmt.Println(generateTrees(3))
}
