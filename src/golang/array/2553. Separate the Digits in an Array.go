package array

func separateDigits(nums []int) []int {
	var res []int
	for _,n := range nums{
		var data []int
		for n > 0{
			m := n % 10
			data = append(data,m)
			n /= 10
		}
		var l int = len(data)
		for i := l - 1;i >= 0;i--{
			res = append(res,data[i])
		}
	}
	return res
}