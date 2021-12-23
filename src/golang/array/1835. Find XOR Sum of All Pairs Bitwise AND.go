package array

//输入：arr1 = [1,2,3], arr2 = [6,5]
//输出：0
//解释：列表 = [1 AND 6, 1 AND 5, 2 AND 6, 2 AND 5, 3 AND 6, 3 AND 5] = [0,1,2,0,2,1] ，
//异或和 = 0 XOR 1 XOR 2 XOR 0 XOR 2 XOR 1 = 0 。

//arr1[i] & arr2[j]
func getXORSum(arr1 []int, arr2 []int) int {
	var l1 int = len(arr1)
	var l2 int = len(arr2)
	var res1 int = 0
	for i := 0; i < l1; i++ {
		res1 ^= arr1[i]
	}
	var res2 int = 0
	for j := 0; j < l2; j++ {
		res2 ^= arr2[j]
	}
	return res1 & res2
}

func getXORSum2(arr1 []int, arr2 []int) int {
	var l1 int = len(arr1)
	var l2 int = len(arr2)
	var res int = 0
	for i := 0; i < 31; i++ {
		cnt1 := 0
		for j := 0; j < l1; j++ {
			if (arr1[j] & (1 << i)) != 0 {
				cnt1++
			}
		}
		cnt2 := 0
		for j := 0; j < l2; j++ {
			if (arr2[j] & (1 << i)) != 0 {
				cnt2++
			}
		}
		if (cnt1|1) == cnt1 && (cnt2|1) == cnt2 {
			res |= 1 << i
		}
	}
	return res
}
