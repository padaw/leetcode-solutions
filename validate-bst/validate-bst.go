package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Val < min || node.Val > max {
		return false
	}
	return walk(node.Left, min, node.Val-1) && walk(node.Right, node.Val+1, max)
}

func isValidBST(root *TreeNode) bool {
	return walk(root, math.MinInt, math.MaxInt)
}
