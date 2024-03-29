package array

//Input: [1,1,1,0,6,12]
//Output: 4
//Explanation: left = [1,1,1,0], right = [6,12]
func PartitionDisjoint(A []int) int {
	var l int = len(A)
	var left []int= make([]int,l)
	var right []int = make([]int,l)
	left[0] = A[0]//store the max number from left side
	right[l - 1] = A[l - 1]//store the min num from right side
	for i := 1;i < l;i++{
		if A[i] > left[i - 1]{
			left[i] = A[i]
		}else{
			left[i] = left[i - 1]
		}
		if A[l - 1 - i] < right[l - i]{
			right[l - 1 - i] = A[l - 1 - i]
		}else{
			right[l - 1 - i] = right[l - i]
		}
	}
	visit := 0
	for visit < l{
		if left[visit] <= right[visit + 1]{//左边最大值 小于 右边最小值
			break
		}else{
			visit++
		}
	}
	return visit + 1
}