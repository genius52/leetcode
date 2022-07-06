package array

//Input: ratings = [1,0,2]
//Output: 5
//Input: ratings = [1,2,2]
//Output: 4
func candy(ratings []int) int {
	var l int = len(ratings)
	var candy_nums []int = make([]int,l)
	for i := 1;i < l;i++{
		if ratings[i] > ratings[i - 1]{
			candy_nums[i] = candy_nums[i - 1] + 1
		}
	}
	var res int = candy_nums[l - 1]
	for i := l - 1;i > 0;i--{
		if ratings[i - 1] > ratings[i]{
			if candy_nums[i - 1] <= candy_nums[i]{
				candy_nums[i - 1] = candy_nums[i] + 1
			}
		}
		res += candy_nums[i - 1]
	}
	return res + l
}