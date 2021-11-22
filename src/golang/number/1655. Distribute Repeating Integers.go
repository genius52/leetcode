package number

import "sort"

//Input: nums = [1,1,2,3], quantity = [2,2]
//Output: false
//Explanation: Although the 0th customer could be given [1,1], the 1st customer cannot be satisfied.
func dp_canDistribute(record []int,quantity []int,l int,pos int)bool{
	if pos >= l{
		return true
	}
	for i := 0;i < len(record);i++{
		if quantity[pos] <= record[i]{
			record[i] -= quantity[pos]
			ret := dp_canDistribute(record,quantity,l,pos + 1)
			if ret{
				return true
			}
			record[i] += quantity[pos]
		}
	}
	return false
}

func CanDistribute(nums []int, quantity []int) bool {
	var cnt [1001]int
	for _,n := range nums{
		cnt[n]++
	}
	var record []int
	for i := 1;i <= 1000;i++{
		if cnt[i] > 0{
			record = append(record,cnt[i])
		}
	}
	//sort.Ints(record)
	sort.Slice(quantity, func(i, j int) bool {
		return quantity[i] >= quantity[j]
	})
	return dp_canDistribute(record,quantity,len(quantity),0)
}