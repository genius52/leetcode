package number

import "sort"

func MinWastedSpace(packages []int, boxes [][]int) int {
	var p_len int = len(packages)
	sort.Ints(packages)
	var l int = len(boxes)
	var res int = 2147483647
	for i := 0; i < l; i++ {
		sort.Ints(boxes[i])
		var cur_len int = len(boxes[i])
		if boxes[i][cur_len-1] < packages[p_len-1] {
			continue
		}
		var cur int = 0
		var low int = 0
		for _, p := range packages {
			var high int = cur_len - 1
			var ret int = low
			for low < high {
				mid := (low + high) / 2
				if boxes[i][mid] < p {
					low = mid + 1
					ret = mid + 1
				} else if boxes[i][mid] == p {
					high = mid
					ret = mid
					break
				} else {
					high = mid
					ret = mid
				}
			}
			cur += boxes[i][ret] - p
			low = ret
			cur %= (1e9 + 7)
			if cur >= res {
				break
			}
		}
		if cur < res {
			res = cur
		}
	}
	if res == 2147483647 {
		return -1
	}
	return res
}
