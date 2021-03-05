package Easy

func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return make([]float64, 0)
    }
    avgs := make([]float64, 0, 10)
    queue := make([]*TreeNode, 1, 10)
    queue[0] = root
    for len(queue) != 0 {
        n := len(queue)
        sum := 0
        for i := 0; i < n; i++ {
            cur := queue[0]
            sum += cur.Val
            queue = queue[1:]
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        avgs = append(avgs, float64(sum)/float64(n))
    }
    return avgs
}
