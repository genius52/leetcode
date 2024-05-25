package array

func IsArraySpecial2(nums []int, queries [][]int) []bool {
	var l1 int = len(nums)
	var l2 int = len(queries)
	var res []bool = make([]bool, l2)
	var prefix []int = make([]int, l1)
	for i := 1; i < l1; i++ {
		if (nums[i-1]^nums[i])&1 == 1 {
			prefix[i] = prefix[i-1] + 1
		} else {
			prefix[i] = prefix[i-1]
		}
	}
	for idx, q := range queries {
		left := q[0]
		right := q[1]
		cnt := prefix[right] - prefix[left]
		if cnt == right-left {
			res[idx] = true
		}
	}
	return res
}
