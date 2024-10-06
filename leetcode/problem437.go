/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    m := make(map[int]int)
    if root == nil {
        return 0
    }
    return count(root, 0, targetSum, m)
}

func count(root *TreeNode, sum int,  targetSum int, m map[int]int) int {
    m1 := copyMap(m)
    ans := 0
    sum += root.Val
    //m1[sum]++
    if sum == targetSum{
        ans++
    }
    ans += m1[sum-targetSum]
    m1[sum]++
	switch {
	case root.Right != nil && root.Left != nil:
        return count(root.Left, sum, targetSum, m1) + count(root.Right, sum, targetSum, m1) + ans
	case root.Left != nil:
        return count(root.Left, sum, targetSum, m1) + ans
	case root.Right != nil:
        return count(root.Right, sum, targetSum, m1) + ans
	default:
		return ans
	}
}

//from m2 to m1
func copyMap(m2 map[int]int) map[int]int {
    m1 := make(map[int]int)
    for k, v := range m2 {
        m1[k] = v
    }
    return m1
}
