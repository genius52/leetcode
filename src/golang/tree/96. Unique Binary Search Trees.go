package tree


//96
//Input: 3
//Output: 5
//Explanation:
//Given n = 3, there are a total of 5 unique BST's:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
func numTrees(n int) int {
	var dp []int = make([]int,n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2;i <= n;i++{
		for j := 0;j < i;j++{
			dp[i] += dp[i - j - 1] * dp[j]
		}
	}
	return dp[n]
}
