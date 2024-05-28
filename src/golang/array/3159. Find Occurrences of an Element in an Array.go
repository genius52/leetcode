package array

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	var l int = len(nums)
	var idx []int
	for i := 0; i < l; i++ {
		if nums[i] == x {
			idx = append(idx, i)
		}
	}
	var l2 int = len(queries)
	var res []int = make([]int, l2)
	for i := 0; i < l2; i++ {
		if queries[i] > len(idx) {
			res[i] = -1
		} else {
			res[i] = idx[queries[i]-1]
		}
	}
	return res
}
