package array

//Input: arr = [2,1,3,4,4]
//Output: 4
//Explanation:
//We can split into two chunks, such as [2, 1], [3, 4, 4].
//However, splitting into [2, 1], [3], [4], [4] is the highest number of chunks possible.

//math magic
//func maxChunksToSorted2(arr []int){
//	var l int = len(arr)
//	var sorted []int = make([]int,l)
//	copy(sorted,arr)
//	sort.Ints(sorted)
//	var sum1,sum2 int = 0,0
//	var res int = 0
//	for i := 0;i < l;i++{
//		sum1 += arr[i]
//		sum2 += sorted[i]
//		if sum1 == sum2{
//			res++
//		}
//	}
//	return res
//}

func MaxChunksToSorted2(arr []int) int {
	var l int = len(arr)
	var dp_rightmin []int = make([]int,l)
	dp_rightmin[l - 1] = arr[l - 1]
	for i := l - 2;i >= 0;i--{
		if arr[i] < dp_rightmin[i + 1]{
			dp_rightmin[i] = arr[i]
		}else{
			dp_rightmin[i] = dp_rightmin[i + 1]
		}
	}
	var res int = 0
	var cur_max int = 0
	for i := 0;i < l - 1;i++{
		if cur_max < arr[i]{
			cur_max = arr[i]
		}
		//if cur_max bigger than right nums,should be merge to one group
		if cur_max <= dp_rightmin[i + 1]{
			res++
		}
	}
	return res + 1
}