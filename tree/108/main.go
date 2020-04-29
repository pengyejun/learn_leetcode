package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findRootVal(nums []int,start int,end int) *TreeNode {
	if start>=end {
		return nil
	}

	mid := (start+end)/2

	root := &TreeNode{Val:nums[mid]}
	root.Left=findRootVal(nums,start,mid)
	root.Right=findRootVal(nums,mid+1,end)

	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return findRootVal(nums, 0, len(nums))
}

func main() {
	
}
