package number

import "math"

//tires = [[2,3],[3,4]], changeTime = 5, numLaps = 4
func MinimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	var l int = len(tires)
	var dp []int = make([]int,numLaps + 1)//dp[i], 跑i圈的最短时间
	for i := 0;i <= numLaps;i++{
		dp[i] = 2147483647
	}
	//var factor int = 0
	var eachround_time int = 2147483647
	for i := 0;i < l;i++{
		if tires[i][0] < eachround_time{
			eachround_time = tires[i][0]
			//factor = tires[i][1]
		}
	}
	//var total int = eachround_time
	var neg bool = false
	for i := 1;i <= numLaps;i++{
		if !neg{
			for k := 0;k < l;k++{
				a1 := tires[k][0]
				an := tires[k][0] * int(math.Pow(float64(tires[k][1]),float64(i - 1)))
				sn := (an * tires[k][1] - a1)/(tires[k][1] - 1)
				if sn > 0 && sn < 2147483647{
					dp[i] = min_int(dp[i],sn)
				}
			}
		}
		if dp[i] < 0{
			neg = true
			dp[i] = 2147483647
		}

		//change_perround := changeTime * (i - 1) + eachround_time * i
		//if change_perround > 0{
		//	dp[i] = change_perround
		//}
		//if total > 0{
		//	dp[i] = min_int(dp[i],total)
		//}
		//eachround_time *= factor
		//total += eachround_time
		for j := 1;j < i;j++{
			dp[i] = min_int(dp[i],dp[j] + dp[i - j] + changeTime)
		}
	}
	return dp[numLaps]
}