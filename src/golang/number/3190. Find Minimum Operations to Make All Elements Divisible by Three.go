package number

func minimumOperations3190(nums []int) int {
	var res int = 0
	for _, n := range nums {
		if n%3 != 0 {
			res++
		}
	}
	return res
}
