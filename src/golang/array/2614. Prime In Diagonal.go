package array

import (
	"math"
	"sort"
)

func is_prime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func diagonalPrime(nums [][]int) int {
	var rows int = len(nums)
	var data []int
	for i := 0; i < rows; i++ {
		data = append(data, nums[i][i])
		data = append(data, nums[rows-1-i][i])
	}
	sort.Ints(data)
	for i := len(data) - 1; i >= 0; i-- {
		if is_prime(data[i]) {
			return data[i]
		}
	}
	return 0
}
