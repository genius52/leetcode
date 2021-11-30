package number

//Input: nums = [3,10,5,25,2,8]
//Output: 28
//Explanation: The maximum result is 5 XOR 25 = 28.
type Trie struct {
	next_0 *Trie
	next_1 *Trie
	val int
}

func FindMaximumXOR(nums []int) int {
	var root Trie
	var l int = len(nums)
	for i := 0;i < l;i++{
		var visit *Trie = &root
		for j := 31;j >= 0;j--{
			if nums[i] | (1 << j) == nums[i]{
				if visit.next_1 == nil{
					visit.next_1 = new(Trie)
				}
				visit = visit.next_1
			}else{
				if visit.next_0 == nil{
					visit.next_0 = new(Trie)
				}
				visit = visit.next_0
			}
		}
		visit.val = nums[i]
	}
	var res int = 0
	for i := 0;i < l;i++{
		var find *Trie = &root
		for j := 31;j >= 0;j--{
			if (nums[i] | (1 << j)) == nums[i]{
				if find.next_0 != nil{
					find = find.next_0
				}else{
					find = find.next_1
				}
			}else{
				if find.next_1 != nil{
					find = find.next_1
				}else{
					find = find.next_0
				}
			}
		}
		res = max_int(res,nums[i] ^ find.val)
	}
	return res
}