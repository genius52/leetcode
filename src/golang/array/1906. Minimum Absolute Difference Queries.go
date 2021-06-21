package array

func MinDifference(nums []int, queries [][]int) []int {
	var l int = len(nums)
	var dp [][101]int = make([][101]int,l + 1)
	for i := 0;i < l;i++{
		for j := 1;j <= 100;j++{
			dp[i + 1][j] = dp[i][j]
		}
		dp[i + 1][nums[i]]++
	}
	var q_len int = len(queries)
	var res []int = make([]int,q_len)
	for i := 0;i < q_len;i++{
		begin := queries[i][0]
		end := queries[i][1]
		var pre int = -1
		var min_diff int = 200
		for i := 1;i <= 100;i++{
			cnt := dp[end + 1][i] - dp[begin][i]
			if cnt != 0{
				if pre != -1{
					if (i - pre) < min_diff{
						min_diff = i - pre
					}
				}
				pre = i
			}
		}
		if min_diff == 200{
			res[i] = -1
		}else{
			res[i] = min_diff
		}
	}
	return res
}