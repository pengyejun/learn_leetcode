package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildtree(inorder []int, inStart, inEnd int, postorder []int, poStart, poEnd int, inorderMap map[int]int) *TreeNode {
	if inStart > inEnd || poStart < poEnd {
		return nil
	}
	index := inorderMap[postorder[poStart]]
	root := &TreeNode{Val:postorder[poStart]}
	numleft := index - inStart
	root.Left = buildtree(inorder, inStart, index - 1, postorder, poEnd + numleft - 1, poEnd, inorderMap)
	root.Right = buildtree(inorder, index + 1, inEnd, postorder, poStart - 1, poEnd + numleft, inorderMap)
	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	return buildtree(inorder, 0, len(inorder) - 1, postorder, len(postorder) - 1, 0, inorderMap)
}

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}
