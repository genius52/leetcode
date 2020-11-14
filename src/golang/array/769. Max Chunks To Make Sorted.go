package array

func maxChunksToSorted(arr []int) int {
	var l int = len(arr)
	var cur_max []int = make([]int,l)
	cur_max[0] = arr[0]
	//var max_num int = math.MinInt32
	var res int = 0
	for i := 1;i < l;i++{
		if arr[i] > cur_max[i-1]{
			cur_max[i] = arr[i]
		}else{
			cur_max[i] = cur_max[i - 1]
		}
	}
	for i := 0;i < l;i++{
		if cur_max[i] == i{
			res++
		}
	}
	return res
}