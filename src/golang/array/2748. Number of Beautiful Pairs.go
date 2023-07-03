package array

import "strconv"

func countBeautifulPairs(nums []int) int {
	var res int = 0
	var l int = len(nums)
	for i := 0; i < l; i++ {
		s1 := strconv.Itoa(nums[i])
		n1 := int(s1[0] - '0')
		for j := i + 1; j < l; j++ {
			s2 := strconv.Itoa(nums[j])
			n2 := int(s2[len(s2)-1] - '0')
			if gcd(n1, n2) == 1 {
				res++
			}
		}
	}
	return res
}
