package number

//Input: nums = [0,1,1,3], maximumBit = 2
//Output: [0,3,2,3]
//Explanation: The queries are answered as follows:
//1st query: nums = [0,1,1,3], k = 0 since 0 XOR 1 XOR 1 XOR 3 XOR 0 = 3.
//2nd query: nums = [0,1,1], k = 3 since 0 XOR 1 XOR 1 XOR 3 = 3.
//3rd query: nums = [0,1], k = 2 since 0 XOR 1 XOR 2 = 3.
//4th query: nums = [0], k = 3 since 0 XOR 3 = 3.
func GetMaximumXor(nums []int, maximumBit int) []int {
	var l int = len(nums)
	var res []int = make([]int,l)
	var max_num int = (1 << maximumBit) - 1
	var total int = 0
	for i := 0;i < l;i++{
		total ^= nums[i]
	}
	var index int = 0
	for l > 0{
		res[index] = max_num ^ total
		total ^= nums[l - 1]
		l--
		index++
	}
	return res
}