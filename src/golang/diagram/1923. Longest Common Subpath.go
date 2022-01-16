package diagram

func LongestCommonSubpath(n int, paths [][]int) int {
	var m int = len(paths)
	var min_len int = 2147483647
	for i := 0;i < m;i++{
		if len(paths[i]) < min_len{
			min_len = len(paths[i])
		}
	}
	var low int = 0
	var high int = min_len
	var BASE int = 100001
	var MOD int = 1e11+7
	var last_low int = -1
	for low < high{
		mid := (low + high + 1)/2
		var total_record map[int]int = make(map[int]int)
		for i := 0;i < m;i++{
			var cur_record map[int]bool = make(map[int]bool)
			hash := 0
			multi := 1
			for j := 0;j < len(paths[i]);j++{
				hash = (hash * BASE + paths[i][j]) % MOD
				if j < mid{
					multi *= BASE
					multi %= MOD
				}else{
					hash = (MOD + hash - multi * paths[i][j - mid] % MOD) % MOD
				}
				if j >= mid - 1{
					cur_record[hash] = true
				}
			}
			for k,_ := range cur_record{
				total_record[k]++
			}
		}
		var find bool = false
		for _,v := range total_record{
			if v == m{
				find = true
				break
			}
		}
		if find{
			low = mid
			if low == last_low{
				break
			}
			last_low = low
		}else{
			high = mid - 1
		}
	}
	return low
}