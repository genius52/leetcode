package array

func CountPairs(nums []int, k int) int {
	var l int = len(nums)
	var res int = 0
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < l;i++{
		if _,ok := record[nums[i]];ok{
			for _,idx := range record[nums[i]]{
				if i * idx % k == 0{
					res++
				}
			}
		}
		record[nums[i]] = append(record[nums[i]],i)
	}
	return res
}