package array

func maxProduct(nums []int) int {
	var first int = 0
	var second int = 0
	for _,n := range nums{
		if n > first{
			second = first
			first = n
		}else if n > second{
			second = n
		}
	}
	return (first - 1) * (second - 1)
}