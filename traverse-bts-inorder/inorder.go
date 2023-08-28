type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, list *[]int) {
    if node == nil {
        return
    }
    walk(node.Left, list)
    *list = append(*list, node.Val)
    walk(node.Right, list)
}


func inorderTraversal(root *TreeNode) []int {
    list := []int{}
    walk(root, &list)

    return list
}
