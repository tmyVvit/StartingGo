package binaryTreeCameras

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从底往上，后序遍历 --》 左 右 中
// 每个节点的状态：
// 0 : 无覆盖
// 1 : 有覆盖
// 2 : 有监控

func minCameraCover(root *TreeNode) (res int) {
	var calCameras func(node *TreeNode) int
	calCameras = func(node *TreeNode) int {
		if node == nil {
			return 1
		}
		left := calCameras(node.Left)
		right := calCameras(node.Right)

		// 左右儿子中存在没被覆盖的  那么当前节点需要监控
		if left == 0 || right == 0 {
			res++
			return 2
		}
		// 左右儿子都被覆盖，那么当前节点就没被覆盖
		if left == 1 && right == 1 {
			return 0
		}
		// 左右儿子中存在至少一个有监控
		if left == 2 || right == 2 {
			return 1
		}
		// 不可能会走到这里
		return -1
	}

	if calCameras(root) == 0 {
		res++
	}
	return res
}
