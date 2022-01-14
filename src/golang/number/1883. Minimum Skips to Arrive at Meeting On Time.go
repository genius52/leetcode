package number

import (
	"math"
)

func dfs_minSkips(dist []int, l int, idx int, speed int, pre_dist int, skip_cnt int, memo map[int]map[int]map[int]float32) float32 {
	if skip_cnt < 0 {
		return 1e9
	}
	if idx == l-1 {
		return float32(dist[idx]+pre_dist) / float32(speed)
	}
	//key := strconv.Itoa(idx) + "|" + strconv.Itoa(skip_cnt) + "|" + strconv.Itoa(pre_dist)
	if _, ok1 := memo[idx]; ok1 {
		if _, ok2 := memo[idx][skip_cnt]; ok2 {
			if _, ok3 := memo[idx][skip_cnt][pre_dist]; ok3 {
				return memo[idx][skip_cnt][pre_dist]
			}
		} else {
			memo[idx][skip_cnt] = make(map[int]float32)
		}
	} else {
		memo[idx] = make(map[int]map[int]float32)
		memo[idx][skip_cnt] = make(map[int]float32)
	}
	mod := (dist[idx] + pre_dist) % speed
	var cost float32 = float32(dist[idx]+pre_dist) / float32(speed)
	var res float32
	if mod == 0 {
		res = cost + dfs_minSkips(dist, l, idx+1, speed, 0, skip_cnt, memo)
	} else {
		//not skip
		have_rest := float32(math.Ceil(float64(cost))) + dfs_minSkips(dist, l, idx+1, speed, 0, skip_cnt, memo)
		//skip
		skip := float32(math.Floor(float64(cost))) + dfs_minSkips(dist, l, idx+1, speed, mod, skip_cnt-1, memo)
		if have_rest < skip {
			res = have_rest
		} else {
			res = skip
		}
	}
	memo[idx][skip_cnt][pre_dist] = res
	return res
}

func MinSkips(dist []int, speed int, hoursBefore int) int {
	var sum int = 0
	var l int = len(dist)
	for i := 0; i < l; i++ {
		sum += dist[i]
	}
	if float32(sum)/float32(speed) > float32(hoursBefore) {
		return -1
	}
	var memo map[int]map[int]map[int]float32 = make(map[int]map[int]map[int]float32)
	var res int = -1
	var low int = 0
	var high int = l - 1
	for low < high {
		mid := (low + high) / 2
		hours := dfs_minSkips(dist, l, 0, speed, 0, mid, memo)
		if hours > float32(hoursBefore) {
			low = mid + 1
			res = mid + 1
		} else {
			high = mid
			res = high
		}
	}
	return res
}
