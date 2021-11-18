package array

import "sort"

//Input: nums = [1,1,2,2,2,3]
//Output: [3,1,1,2,2,2]
func frequencySort(nums []int) []int{
	var cnt [201]int
	for _,n := range nums{
		cnt[n + 100]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if cnt[nums[i] + 100] == cnt[nums[j] + 100]{
			return nums[i] > nums[j]
		}else{
			return cnt[nums[i] + 100] < cnt[nums[j] + 100]
		}
	})
	return nums
}

func FrequencySort(nums []int) []int {
	var record map[int]int = make(map[int]int)
	var l int = len(nums)
	for i := 0;i < l;i++{
		record[nums[i]]++
	}
	var freq_record map[int][]int = make(map[int][]int)
	for n,r := range record{
		freq_record[r] = append(freq_record[r],n)
		sort.Ints(freq_record[r])
	}
	var freqs []int
	for freq,_ := range freq_record{
		freqs = append(freqs,freq)
	}
	sort.Ints(freqs)
	var res []int = make([]int,l)
	var index int = 0
	for _,freq := range freqs{
		num_arr := freq_record[freq]
		var arr_len int = len(num_arr)
		for i := arr_len - 1;i >= 0;i--{
			for j := 0;j < freq;j++{
				res[index] = freq_record[freq][i]
				index++
			}
		}
	}
	return res
}