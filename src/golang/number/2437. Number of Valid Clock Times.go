package number

import "strconv"

//00:00--23:59
func dfs_countTime(time string,cur string,idx int)int{
	if idx == 5{
		hours,_ := strconv.Atoi(cur[0:2])
		if hours >= 24{
			return 0
		}
		minutes,_ := strconv.Atoi(cur[3:])
		if minutes >= 60{
			return 0
		}
		return 1
	}
	if idx == 2{
		return dfs_countTime(time,cur + ":",idx + 1)
	}
	if time[idx] == '?'{
		var res int = 0
		for i := 0;i <= 9;i++{
			res += dfs_countTime(time,cur + strconv.Itoa(i),idx + 1)
		}
		return res
	}else{
		return dfs_countTime(time,cur + string(time[idx]),idx + 1)
	}
}

func countTime(time string) int {
	return dfs_countTime(time,"",0)
}