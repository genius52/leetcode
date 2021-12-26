package array

func getDistances(arr []int) []int64 {
	var l int = len(arr)
	var ret []int64 = make([]int64, l)
	var record map[int][]int = make(map[int][]int)
	var prefix map[int][]int64 = make(map[int][]int64)
	for i := 0; i < l; i++ {
		if _, ok := record[arr[i]]; !ok {
			record[arr[i]] = []int{i}
			prefix[arr[i]] = []int64{0}
		} else {
			var cur_len int = len(prefix[arr[i]])
			record[arr[i]] = append(record[arr[i]], i)
			prefix[arr[i]] = append(prefix[arr[i]], (int64(i)-int64(record[arr[i]][cur_len-1]))*int64(cur_len)+prefix[arr[i]][cur_len-1])
		}
	}
	var suffix map[int][]int64 = make(map[int][]int64)
	for key, vals := range record {
		var cur_len int = len(vals)
		if cur_len == 1 {
			suffix[key] = []int64{0}
		} else if cur_len > 1 {
			suffix[key] = make([]int64, cur_len)
			suffix[key][cur_len-1] = 0
			for i := cur_len - 2; i >= 0; i-- {
				suffix[key][i] = (int64(vals[i+1])-int64(vals[i]))*int64(cur_len-i-1) + suffix[key][i+1]
			}
		}
	}
	for key, vals := range record {
		var cur_len int = len(vals)
		for i := 0; i < cur_len; i++ {
			ret[vals[i]] = prefix[key][i] + suffix[key][i]
		}
	}
	return ret
}
