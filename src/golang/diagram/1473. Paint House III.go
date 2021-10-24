package diagram

//输入：houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
//输出：9
//解释：房子涂色方案为 [1,2,2,1,1]
//此方案包含 target = 3 个街区，分别是 [{1}, {2,2}, {1,1}]。
//涂色的总花费为 (1 + 1 + 1 + 1 + 5) = 9。
func dp_minCost(houses []int,pos int,cost [][]int,m int,n int,target int,memo *[][][]int,last_color int)int{
	if pos >= m{
		if target == 0{
			return 0
		}
		return -1
	}
	if target < 0{
		return -1
	}
	if (*memo)[pos][target][last_color] != 0{
		return (*memo)[pos][target][last_color]
	}
	if houses[pos] != 0{
		next_target := target
		if houses[pos] != last_color{
			next_target--
		}
		return dp_minCost(houses,pos + 1,cost,m,n,next_target,memo,houses[pos])
	}
	var res int = -1
	for i := 1;i <= n;i++{
		//houses[pos] = i
		next_target := target
		if i != last_color{
			next_target--
		}
		cur :=  dp_minCost(houses,pos + 1,cost,m,n,next_target,memo,i)
		if cur != -1{
			if res == -1 || cur + cost[pos][i - 1] < res{
				res = cur + cost[pos][i - 1]
			}
		}
		//houses[pos] = 0
	}
	(*memo)[pos][target][last_color] = res
	return res
}

func MinCost1473(houses []int, cost [][]int, m int, n int, target int) int {
	var memo [][][]int = make([][][]int,m)//memo[pos][]
	for i := 0;i < m;i++{
		memo[i] = make([][]int,target + 1)
		for j := 0;j <= target;j++{
			memo[i][j] = make([]int,n + 1)
		}
	}
	return dp_minCost(houses,0,cost,m,n,target,&memo,0)
}