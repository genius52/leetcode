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

func is_prime2(n int) {
	var prime []bool = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}
	for i := 2; i <= n; i++ {
		if prime[i] {
			for j := i * i; j <= n; j += i { //i * i 而不是i + i的原因是为了更高效，减少重复
				prime[j] = false
			}
		}
	}
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
