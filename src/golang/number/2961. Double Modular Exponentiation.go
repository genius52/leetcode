package number

// ((aibi % 10)ci) % mi == target
// (a⋅b)mod m=((a mod m)⋅(b mod m))mod m
func getGoodIndices(variables [][]int, target int) []int {
	var res []int
	for i := 0; i < len(variables); i++ {
		var a int64 = int64(variables[i][0])
		var b int64 = int64(variables[i][1])
		var c int64 = int64(variables[i][2])
		var m int64 = int64(variables[i][3])
		ret := int(myPow2(myPow2(a, b, 10), c, m))
		if ret == target {
			res = append(res, i)
		}
	}
	return res
}

func myPow2(x int64, n int64, mod int64) int64 {
	if n <= 0 {
		return 1
	}
	var res int64 = 1
	res = myPow2(x, n/2, mod)
	res = res * res % int64(mod)
	if n%2 == 1 {
		res = res * int64(x) % int64(mod)
	}
	return res
}
