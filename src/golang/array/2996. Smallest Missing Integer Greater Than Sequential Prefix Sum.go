package array

func MissingInteger(nums []int) int {
	var l int = len(nums)
	if l == 1 {
		return nums[0]
	}
	var sum int = nums[0]
	var i int = 1
	for ; i < l; i++ {
		if nums[i] != nums[i-1]+1 {
			break
		}
		sum += nums[i]
	}
	if i == l {
		return sum
	}
	var visited map[int]bool = make(map[int]bool)
	for i := 0; i < l; i++ {
		visited[nums[i]] = true
	}
	for true {
		if _, ok := visited[sum]; !ok {
			return sum
		}
		sum++
	}
	return sum
}
