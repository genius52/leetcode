package tree

type TreeAncestor struct {
	//2 ** 16 = 65536 > 50000
	dp [][16]int //dp[i][j]: 节点i的 2^j 祖先， dp[i][j] = -1表示节点i没有第2^j个祖先
}


func Constructor1483(n int, parent []int) TreeAncestor {
	var obj TreeAncestor
	obj.dp = make([][16]int,n)
	for i := 0;i < n;i++{
		for j := 0;j < 16;j++ {
			obj.dp[i][j] = -1
		}
	}
	for i := 0;i < n;i++{
		obj.dp[i][0] = parent[i]
	}
	for i := 1;i < 16;i++{
		for j := 0;j < n;j++{
			if obj.dp[j][i - 1] != -1{
				obj.dp[j][i] = obj.dp[obj.dp[j][i - 1]][i - 1]
			}
		}
	}
	return obj
}


func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	var find_node int = node
	for i := 0;i < 16;i++{
		if ((1 << i) | k) == k{
			if this.dp[find_node][i] != -1{
				find_node = this.dp[find_node][i]
			}else{
				return -1
			}
		}
	}
	return find_node
}