package array

//输入：arr = [2,3,1,6,7]
//输出：4
//解释：满足题意的三元组分别是 (0,1,2), (0,2,2), (2,3,4) 以及 (2,4,4)
func CountTriplets(arr []int) int {
	var l int = len(arr)
	var prefix []int = make([]int,l + 1)
	for i := 0;i < l;i++{
		prefix[i + 1] = arr[i] ^ prefix[i]
	}
	var res int = 0
	for i := 0;i < l - 1;i++{
		for j := i + 1;j < l;j++{
			for k := j;k < l;k++{
				a := prefix[i] ^ prefix[j]
				b := prefix[j] ^ prefix[k + 1]
				if a == b {
					res++
				}
			}
		}
	}
	return res
}