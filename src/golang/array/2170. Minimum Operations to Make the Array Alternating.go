package array

func minimumOperations(nums []int) int {
	var even_record map[int]int = make(map[int]int)
	var odd_record map[int]int = make(map[int]int)
	var l int = len(nums)
	for i := 0;i < l;i++{
		if i | 1 == i{
			odd_record[nums[i]]++
		}else{
			even_record[nums[i]]++
		}
	}
	var first_even_num int = 0
	var first_even_freq int = 0
	//var second_even_num int = 0
	var second_even_freq int = 0
	for k,v := range even_record{
		if v > first_even_freq{
			if first_even_freq > second_even_freq{
				//second_even_num = first_even_num
				second_even_freq = first_even_freq
			}
			first_even_num = k
			first_even_freq = v
		}else{
			if v > second_even_freq{
				//second_even_num = k
				second_even_freq = v
			}
		}
	}
	var first_odd_num int = 0
	var first_odd_freq int = 0
	//var second_odd_num int = 0
	var second_odd_freq int = 0
	for k,v := range odd_record{
		if v > first_odd_freq{
			if first_odd_freq > second_odd_freq{
				//second_odd_num = first_odd_num
				second_odd_freq = first_odd_freq
			}
			first_odd_num = k
			first_odd_freq = v
		}else{
			if v > second_odd_freq{
				//second_odd_num = k
				second_odd_freq = v
			}
		}
	}
	var res int = 0
	if first_even_num != first_odd_num{
		res = (l + 1)/2 - first_even_freq + l/2 - first_odd_freq
	}else{
		res = min_int((l + 1)/2 - first_even_freq + l/2 - second_odd_freq,(l + 1)/2 - second_even_freq + l/2 - first_odd_freq)
	}
	return res
}