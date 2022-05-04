package number

//Input: nums = [1,2,5,9], threshold = 6
//Output: 5
//Explanation: We can get a sum to 17 (1+2+5+9) if the divisor is 1.
//If the divisor is 4 we can get a sum to 7 (1+1+2+3) and if the divisor is 5 the sum will be 5 (1+1+1+2).
func cal_smallestDivisor(nums []int,div int)int{
	var res int = 0
	for _,n := range nums{
		res += n / div
		if n % div != 0{
			res++
		}
	}
	return res
}

func smallestDivisor(nums []int, threshold int) int {
	var low int = 1
	var l int = len(nums)
	var high int = 0
	for i := 0;i < l;i++{
		if nums[i] > high{
			high = nums[i]
		}
	}
	for low < high{//after loop, high is the biggest
		var mid = (low + high)/2
		ret := cal_smallestDivisor(nums,mid)
		if ret <= threshold{
			high = mid
		}
		if ret > threshold{
			low = mid + 1
		}
	}
	return low
}