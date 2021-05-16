package Hard

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const (
	NotMonitored      = 0
	MonitoredNoCam    = 1
	MonitoredWithCam  = 2
)

func minCameraCover(root *TreeNode) int {
	cameras := 0
	top := dfs(root, &cameras)

	if top == NotMonitored {
		return cameras + 1
	} else {
		return cameras
	}
}

func dfs(root *TreeNode, cameras *int) int {
	if root == nil {
		return MonitoredNoCam
	}

	left, right := dfs(root.Left, cameras), dfs(root.Right, cameras)
	if left == MonitoredNoCam && right == MonitoredNoCam {
		return NotMonitored
	} else if left == NotMonitored || right == NotMonitored {
		*cameras++
		return MonitoredWithCam
	} else {
		return MonitoredNoCam
	}
}
