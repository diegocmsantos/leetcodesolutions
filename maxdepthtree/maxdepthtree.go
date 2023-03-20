package maxdepthtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth, rDepth := 1, 1
	currL, currR := root.Left, root.Right
	for currL != nil || currR != nil {
		if currL.Left != nil {
			lDepth++
			currL = currL.Left
		}
		if currR.Right != nil {
			rDepth++
			currR = currR.Right
		}
	}
	return max(lDepth, rDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
