package number

import "sort"

func CountExcellentPairs(nums []int, k int) int64 {
	var record map[int]int = make(map[int]int)
	var l int = len(nums)
	for i := 0;i < l;i++{
		n := nums[i]
		if _,ok := record[n];!ok{
			var one_cnt int = 0
			for n > 0{
				if n | 1 == n{
					one_cnt++
				}
				n >>= 1
			}
			record[nums[i]] = one_cnt
		}
	}
	var cnt_num [][2]int
	for k,v := range record{
		cnt_num = append(cnt_num,[2]int{v,k})
	}
	sort.Slice(cnt_num, func(i, j int) bool {
		return cnt_num[i][0] < cnt_num[j][0]
	})
	var res int64 = 0
	var l2 int = len(cnt_num)
	for i := 0;i < l2;i++{
		cnt1 := cnt_num[i][0]
		num1 := cnt_num[i][1]
		if cnt1 * 2 >= k{
			res += int64(l2 - i - 1) * 2 + 1
		}else{
			find_idx := sort.Search(l2, func(j int) bool {
				return cnt1 + cnt_num[j][0] >= k
			})
			if find_idx < l2 && find_idx >= i{
				res += int64(l2 - find_idx) * 2
				if num1 == cnt_num[find_idx][1]{
					res--
				}
			}
		}
	}
	return res
}