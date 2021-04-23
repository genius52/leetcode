package array

func count_length(nums []int,l int,start int,visited []bool)int{
	var steps int = 0
	var index int = start
	for !visited[index]{
		visited[index] = true
		index = nums[index]
		steps++
	}
	return steps
}

func arrayNesting(nums []int) int {
	var l int = len(nums)
	var visited []bool = make([]bool,l)
	var max_len int = 0
	for i := 0;i < l;i++{
		if visited[i]{
			continue
		}
		var cur_len int = count_length(nums,l,i,visited)
		if cur_len > max_len{
			max_len = cur_len
		}
	}
	return max_len
}