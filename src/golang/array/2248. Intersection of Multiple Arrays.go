package array

func intersection2248(nums [][]int) []int {
	var l int = len(nums)
	var record [1001]int
	for i := 0;i < l;i++{
		for _,n := range nums[i]{
			record[n]++
		}
	}
	var res []int
	for i := 1;i <= 1000;i++{
		if record[i] == l{
			res = append(res,i)
		}
	}
	return res
}