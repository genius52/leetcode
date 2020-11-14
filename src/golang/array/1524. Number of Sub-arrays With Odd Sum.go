package array

//Input: arr = [1,3,5]
//Output: 4
//Explanation: All sub-arrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
//All sub-arrays sum are [1,4,9,3,8,5].
//Odd sums are [1,9,3,5] so the answer is 4.

func NumOfSubarrays(arr []int) int {
	var odd []int
	var even []int
	for _,n := range arr{
		if(n % 2 == 1){
			odd = append(odd,n)
		}else{
			even = append(even,n)
		}
	}
	if(len(odd) == 0){
		return 0
	}
	//var mod int = 10e9 + 7
	var res int = 0
	return res
}

//数学方法计算排列数(从n中取m个数)
//func mathPailie(n int, m int) int {
//	return jieCheng(n) / jieCheng(n-m)
//}

//数学方法计算组合数(从n中取m个数)
func mathZuhe(n int, m int) int {
	return jieCheng(n) / (jieCheng(n-m) * jieCheng(m))
}

//阶乘
func jieCheng(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
