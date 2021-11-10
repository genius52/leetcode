package diagram

//Input: locations = [4,3,1], start = 1, finish = 0, fuel = 6
//Output: 5
//Explanation: The following are all possible routes:
//1 -> 0, used fuel = 1
//1 -> 2 -> 0, used fuel = 5
//1 -> 2 -> 1 -> 0, used fuel = 5
//1 -> 0 -> 1 -> 0, used fuel = 3
//1 -> 0 -> 1 -> 0 -> 1 -> 0, used fuel = 5
func dfs_countRoutes(locations []int,l int,cur int,finish int,fuel int,memo [][]int)int{
	if fuel < 0{
		return 0
	}
	if memo[cur][fuel] != -1{
		return memo[cur][fuel]
	}
	var res int = 0
	if cur == finish{
		res++
	}
	for i := 0;i < l;i++{
		if i == cur{
			continue
		}
		left := fuel - abs_int(locations[i] - locations[cur])
		if left >= 0{
			res += dfs_countRoutes(locations,l,i,finish,left,memo)
			res %= (1e9 + 7)
		}
	}
	memo[cur][fuel] = res
	return res
}

func CountRoutes(locations []int, start int, finish int, fuel int) int {
	var l int = len(locations)
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,fuel + 1)
		for j := 0;j <= fuel;j++{
			memo[i][j] = -1
		}
	}
	return dfs_countRoutes(locations,l,start,finish,fuel,memo)
}