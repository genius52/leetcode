package number

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0]{
			return properties[i][1] > properties[j][1]
		}
		return properties[i][0] < properties[j][0]
	})
	var res int = 0
	var l int = len(properties)
	var max_def int = -2147483648
	for i := l - 1;i >= 0;i--{
		if properties[i][1] < max_def{
			res++
		}
		if properties[i][1] > max_def{
			max_def = properties[i][1]
		}
	}
	return res
}