package diagram

//Input: cost = [[1, 3, 5], [4, 1, 1], [1, 5, 3]]
//Output: 4
//Explanation: The optimal way of connecting the groups is:
//1--A
//2--B
//2--C
//3--A
//This results in a total cost of 4.
//size1 >= size2

//问题等价于: 在一个矩阵中选取一些值, 满足矩阵的每一行和每一列都至少有一个元素被选中,
//同时选中元素的总和最小 (此矩阵就是 cost 矩阵).
func dfs_connectTwoGroups(cost [][]int,min_cost2 []int,size1 int,size2 int,pos1 int,pos2 int,status int,target int,memo [][]int)int{
	var res int = 2147483647
	if memo[pos1][status] != 0{
		return memo[pos1][status]
	}
	if pos1 < size1{
		for i := 0;i < size2;i++{
			cur := cost[pos1][i] + dfs_connectTwoGroups(cost,min_cost2,size1,size2,pos1 + 1,pos2,status | 1 << i,target,memo)
			res = min_int(res,cur)
		}
	}else{
		var cost int = 0
		for i := 0;i < size2;i++{
			if (status | 1 << i) == status{
				continue
			}
			cost += min_cost2[i]
		}
		res = cost
	}
	memo[pos1][status] = res
	return res
}

func ConnectTwoGroups(cost [][]int) int {
	var size1 int = len(cost)
	var size2 int = len(cost[0])
	var min_cost2 []int = make([]int,size2)
	for j := 0;j < size2;j++{
		var min_val int = 2147483647
		for i := 0;i < size1;i++{
			min_val = min_int(min_val,cost[i][j])
		}
		min_cost2[j] = min_val
	}
	var memo [][]int = make([][]int,size1 + 1)
	for i := 0;i <= size1;i++{
		memo[i] = make([]int,1 << size2)
	}
	return dfs_connectTwoGroups(cost,min_cost2,size1,size2,0,0,0,1 << size2 - 1,memo)
}