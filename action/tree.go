package action

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitTree() TreeNode {
	return TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Right: &TreeNode{Val: 7}, Left: &TreeNode{Val: 17}},
	}
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		temValueList := make([]int, 0)
		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:]
			temValueList = append(temValueList, first.Val)
			if first.Left != nil {
				tmpQueue = append(tmpQueue, first.Left)
			}
			if first.Right != nil {
				tmpQueue = append(tmpQueue, first.Right)
			}

		}
		queue = queue[0:0]
		queue = append(queue, tmpQueue...)
		res = append(res, temValueList)
	}
	return res
}
