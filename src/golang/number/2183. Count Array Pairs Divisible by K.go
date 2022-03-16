package number

//func gcd(a int, b int) int {
//	if b == 0 {
//		return a
//	}
//	return gcd(b, a%b)
//}

func countPairs2183(nums []int, k int) int64 {
	var record map[int]int64 = make(map[int]int64)
	var res int64
	for _,n := range nums{
		g := gcd(n,k)
		for n2,cnt := range record{
			if g * n2 % k == 0{
				res += cnt
			}
		}
		record[g]++
	}
	return res
}