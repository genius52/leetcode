package array

func findMatrix(nums []int) [][]int {
	var record map[int]int = make(map[int]int)
	var l int = len(nums)
	var rows int = 0
	for i := 0; i < l; i++ {
		record[nums[i]]++
		if record[nums[i]] > rows {
			rows = record[nums[i]]
		}
	}
	var res [][]int = make([][]int, rows)
	for k, v := range record {
		for i := 0; i < v; i++ {
			res[i] = append(res[i], k)
		}
	}
	return res
}
