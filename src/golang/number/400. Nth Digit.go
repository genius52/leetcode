package number

import (
	"math"
	"strconv"
)

//Input:
//11
//Output:
//0
func FindNthDigit(n int) int {
	if(n <= 9){
		return n
	}
	var num_cnt int = 9
	var each_num_len int = 1
	var lengths int = num_cnt * each_num_len
	var total_len int = lengths
	for n / total_len > 0{
		num_cnt *= 10
		each_num_len++
		lengths = num_cnt * each_num_len
		total_len += lengths
	}
	var diff int = n - (total_len - lengths)
	var diff_num int = int(math.Pow(10.0,float64(each_num_len - 1))) - 1 +  diff / each_num_len
	var bit_offset int = diff % each_num_len
	if(bit_offset == 0){
		num_str := strconv.Itoa(diff_num)
		res := int(num_str[each_num_len - 1] - '0')
		return res
	}
	num_str := strconv.Itoa(diff_num + 1)
	res := int(num_str[bit_offset - 1] - '0')
	return res
}
