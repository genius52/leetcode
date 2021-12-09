package number

//queries[i] = [favoriteType, favoriteDay, dailyCap]
//糖果类型索引，天数，每天最多吃几颗
func CanEat(candiesCount []int, queries [][]int) []bool {
	var l int = len(candiesCount)
	var prefix []int64 = make([]int64, l+1)
	for i := 0; i < l; i++ {
		prefix[i+1] = int64(candiesCount[i]) + prefix[i]
	}
	var l2 int = len(queries)
	var res []bool = make([]bool, l2)
	for i := 0; i < l2; i++ {
		min_candies := queries[i][1] + 1
		max_candies := (queries[i][1] + 1) * queries[i][2]
		min_target_candy := prefix[queries[i][0]] + 1
		max_target_candy := prefix[queries[i][0]+1]
		if int64(min_candies) > max_target_candy || int64(max_candies) < min_target_candy {
			res[i] = false
		} else {
			res[i] = true
		}
	}
	return res
}
