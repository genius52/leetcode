package array

//输入：nums = [1,0,0,0,0,0,1,1], k = 3
//输出：5
//解释：通过 5 次操作，最左边的 1 可以移到右边直到 nums 变为 [0,0,0,0,0,1,1,1] 。
//TO DO
func MinMoves(nums []int, k int) int {
	var l int = len(nums)
	var record []int//所有1的位置
	for i := 0;i < l;i++{
		if nums[i] == 1{
			record = append(record,i)
		}
	}
	var l2 int = len(record)
	if l2 < k{
		return -1
	}
	//var prefix_one_cnt []int = make([]int,l2 + 1)
	//for i := 0;i < l2;i++{
	//	prefix_one_cnt[i + 1] = prefix_one_cnt[i] + 1 //
	//}
	var res int = 2147483647
	for i := 0;i + k <= l2;i++{
		mid_pos1 := (record[i + (k - 1 )/ 2] + record[i + k/2] )/2
		cnt := k/2
		var left_sum int = 0
		for j := i;record[j] <= mid_pos1;j++{
			left_sum += mid_pos1 - record[j] - cnt
			cnt--
		}
		var right_sum int = 0
		cnt = k/2
		mid_pos2 := mid_pos1
		if k | 1 != k{
			mid_pos2++
		}
		for j := i + k - 1;record[j] >= mid_pos2;j--{
			right_sum += record[j] - mid_pos2 - cnt
			cnt--
		}
		res = min_int(res, left_sum + right_sum)
	}
	return res
}