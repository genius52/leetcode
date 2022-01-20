package array

func dp_colorTheGrid(m int, n int,r int,c int,data [][]int,memo [][1024]int)int{
	if c >= n{
		return 1
	}
	if r == m{
		return dp_colorTheGrid(m,n,0,c + 1,data,memo)
	}
	var pre int = 0
	if c != 0{
		pre = data[r][c - 1]
	}
	var up int = 0
	if r > 0{
		up = data[r - 1][c]
	}
	var pre_mask int = 0
	if r == 0{
		if c > 0{
			for i := 0;i < m;i++{
				pre_mask *= 4
				pre_mask += data[i][c - 1]
			}
		}
		if memo[c][pre_mask] != 0{
			return memo[c][pre_mask]
		}
	}


	var MOD int = 1e9 + 7
	var res int = 0
	for i := 1;i <= 3;i++{
		if i != pre && i != up{
			data[r][c] = i
			res += dp_colorTheGrid(m,n,r + 1,c,data,memo)
			res %= MOD
		}
	}
	if r == 0{
		memo[c][pre_mask] = res
	}
	return res
}

func ColorTheGrid(m int, n int) int {
	var data [][]int = make([][]int,m + 1)
	for i := 0;i <= m;i++{
		data[i] = make([]int,n + 1)
	}
	var memo [][1024]int = make([][1024]int,n + 1)
	return dp_colorTheGrid(m,n,0,0,data,memo)
}