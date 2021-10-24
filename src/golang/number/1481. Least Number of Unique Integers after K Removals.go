package number

import "sort"

//移除出现次数最少的数字
func findLeastNumOfUniqueInts(arr []int, k int) int {
	var record map[int]int = make(map[int]int)
	for _,n := range arr{
		record[n]++
	}
	var cnt []int
	for _,val := range record{
		cnt = append(cnt,val)
	}
	sort.Ints(cnt)
	var idx int = 0
	var l int = len(cnt)
	for k > 0 && idx < l{
		if cnt[idx] <= k{
			k -= cnt[idx]
			idx++
		}else{
			break
		}
	}
	return l - idx
}