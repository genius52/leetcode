package array

import "math"

func GetMaxLen(nums []int) int {
	nums = append(nums,0)
	var l int = len(nums)
	var begin int = 0
	var end int = 0
	var max_length int = 0
	var neg_cnt int = 0
	for end < l{
		if nums[end] < 0{
			neg_cnt++
		}else if nums[end] == 0{
			if neg_cnt % 2 == 0{
				max_length = int(math.Max(float64(max_length),float64(end - begin)))
			}else{
				for begin < end{
					if nums[begin] < 0{
						max_length = int(math.Max(float64(max_length),float64(end - begin - 1)))
						break
					}
					begin++
				}
			}
			neg_cnt = 0
			begin = end + 1
			for begin < l{
				if nums[begin] != 0{
					break
				}
				begin++
			}
			end = begin
			continue
		}
		if neg_cnt % 2 == 0{
			max_length = int(math.Max(float64(max_length),float64(end - begin + 1)))
		}
		end++
	}
	return max_length
}