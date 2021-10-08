package number

//Input: n = 13
//Output: 4
//Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
//[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups with largest size.
func CountLargestGroup(n int) int {
	var record map[int]int = make(map[int]int)
	var max_cnt int = 0
	for i := 1;i <= n;i++{
		var sum int = 0
		var num int = i
		for num > 0{
			sum += num % 10
			num = num / 10
		}
		record[sum]++
		if record[sum] > max_cnt{
			max_cnt = record[sum]
		}
	}
	var res int = 0
	for _,cnt := range record{
		if cnt == max_cnt{
			res++
		}
	}
	return res
}