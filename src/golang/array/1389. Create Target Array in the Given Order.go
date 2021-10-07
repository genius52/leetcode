package array

func CreateTargetArray(nums []int, index []int) []int {
	var l int = len(nums)
	var res []int
	for i := 0;i < l;i++{
		rear := append([]int{},res[index[i]:]...)
		res = append(res[:index[i]],nums[i])
		res = append(res,rear...)
	}
	return res
}