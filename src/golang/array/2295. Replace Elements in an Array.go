package array

func arrayChange(nums []int, operations [][]int) []int {
	var l int = len(nums)
	var num_idx map[int][]int = make(map[int][]int)
	for i := 0;i < l;i++{
		if _,ok := num_idx[nums[i]];ok{
			num_idx[nums[i]] = append(num_idx[nums[i]],i)
		}else{
			num_idx[nums[i]] = []int{i}
		}
	}
	for _,ope := range operations{
		for _,idx := range num_idx[ope[0]]{
			nums[idx] = ope[1]
			if _,ok := num_idx[ope[1]];ok{
				num_idx[ope[1]] = append(num_idx[ope[1]],idx)
			}else{
				num_idx[ope[1]] = []int{idx}
			}
		}
		delete(num_idx,ope[0])
	}
	return nums
}