package array

//The number at the ith position is divisible by i.
//i is divisible by the number at the ith position.
func check_nums(nums []int)bool{
	for i,n := range nums{
		if n % (i + 1) == 0 || (i + 1) % n == 0{
			continue
		}
		return false
	}
	return true
}

func perm(nums []int,l int,pos int)int{
	if pos == l{
		if check_nums(nums){
			return 1
		}
		return 0
	}
	var total int = 0
	for i := pos;i < l;i++{
		if (nums[i] % (pos + 1) == 0 || (pos + 1) % nums[i] == 0) {
			nums[i],nums[pos] = nums[pos],nums[i]
			total += perm(nums,l,pos + 1)
			nums[i],nums[pos] = nums[pos],nums[i]
		}
	}
	return total
}

func CountArrangement(N int) int {
	var nums []int = make([]int,N)
	for i := 0;i < N;i++{
		nums[i] = i + 1
	}
	return perm(nums,N,0)
}