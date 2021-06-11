package number

import "math"

//Person A will NOT friend request person B (B != A) if any of the following conditions are true:
//age[B] <= 0.5 * age[A] + 7
//age[B] > age[A]
//age[B] > 100 && age[A] < 100

func numFriendRequests(ages []int) int {
	var cnt [121]int
	for _,age := range ages{
		cnt[age]++
	}
	var prefix [121]int
	for i := 1;i <= 120;i++{
		prefix[i] = prefix[i - 1] + cnt[i]
	}
	var res int = 0
	for i := 120;i >= 15;i--{
		if cnt[i] == 0{
			continue
		}
		//find younger than i and bigger than 0.5 * i + 7
		//res += cnt[i] * (cnt[i] - 1)
		//向下取整 floor
		low := int(math.Floor(float64(i/2) + 7))
		cur_cnt := prefix[i] - prefix[low]
		res += cur_cnt * cnt[i] - cnt[i]
	}
	return res
}