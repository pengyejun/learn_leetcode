package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildtree(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int, inoderMap map[int]int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	root := &TreeNode{Val:preorder[preStart]}
	inRoot := inoderMap[root.Val]
	numsLeft := inRoot - inStart
	root.Left = buildtree(preorder, preStart + 1, preStart + numsLeft, inorder, inStart, inRoot - 1, inoderMap)
	root.Right = buildtree(preorder, preStart + numsLeft + 1, preEnd, inorder, inRoot + 1, inEnd, inoderMap)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i, v := range inorder {
		m[v] = i
	}
	return buildtree(preorder, 0, len(preorder) - 1, inorder, 0, len(inorder) - 1, m)
}

func main() {
	preorder := []int{3, 9, 8, 10, 20, 15, 7}
	inorder := []int{8, 9, 10, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}
