package number

import (
	"sort"
	"strconv"
)

func SmallestTrimmedNumbers(nums []string, queries [][]int) []int {
	var l int = len(queries)
	var res []int = make([]int,l)
	var l2 int = len(nums)
	var l3 int = len(nums[0])
	for i := 0;i < l;i++{
		var data [][2]string = make([][2]string,l2)
		for j := 0;j < l2;j++{
			data[j] = [2]string{nums[j][max_int(0,l3 - queries[i][1]):],strconv.Itoa(j)}
		}
		sort.Slice(data, func(m, n int) bool {
			var cur_len int = len(data[m][0])
			for k := 0;k < cur_len;k++{
				if data[m][0][k] != data[n][0][k]{
					return data[m][0][k] < data[n][0][k]
				}
			}
			idx1,_ := strconv.Atoi(data[m][1])
			idx2,_ := strconv.Atoi(data[n][1])
			return idx1 < idx2
		})
		idx,_ := strconv.Atoi(data[queries[i][0] - 1][1])
		res[i] = idx
	}
	return res
}