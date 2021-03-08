package number

import "math"

func recursive_checkPowersOfThree(n int64,visited []bool,memo []bool)bool{
	if n == 0{
		return true
	}
	if memo[n]{
		return false
	}
	var res bool = false
	for i := 0;i < 32;i++{
		var choose int64 = int64(math.Pow(3,float64(i)))
		if choose > n{
			break
		}
		if visited[i]{
			continue
		}
		visited[i] = true
		res := recursive_checkPowersOfThree(n - choose,visited,memo)
		if res{
			return true
		}
		visited[i] = false
	}
	memo[n] = true
	return res
}

func CheckPowersOfThree(n int) bool {
	var visited []bool = make([]bool,32)
	var memo []bool = make([]bool,n + 1)
	return recursive_checkPowersOfThree(int64(n),visited,memo)
}