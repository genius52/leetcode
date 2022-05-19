package array

import "sort"

func MaximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	var l int = len(tiles)
	var res int = 0
	var left int = 0
	var right int = 0
	var cnt int = 0
	for left < l && right < l{
		for right < l && (tiles[right][1] - tiles[left][0] + 1) <= carpetLen{
			cnt += tiles[right][1] - tiles[right][0] + 1
			right++
		}
		if right < l{
			res = max_int(res,cnt + max_int(0,tiles[left][0] + carpetLen - tiles[right][0]))
		}else{
			res = max_int(res,cnt)
		}
		cnt -= tiles[left][1] - tiles[left][0] + 1
		left++
	}
	return res
}