package array

//Here is my idea: instead of calculating area by height*width, we can think it in a cumulative way. In other words, sum water amount of each bin(width=1).
//Search from left to right and maintain a max height of left and right separately, which is like a one-side wall of partial container.
// Fix the higher one and flow water from the lower part. For example, if current height of left is lower, we fill water in the left bin.
// Until left meets right, we filled the whole container
//Two pointer
func trap(height []int) int {
	var cap int
	var left_max int = 0
	var right_max int = 0
	var low int = 0
	var high int = len(height) - 1
	for low <= high{
		if(height[low] < height[high]){
			if(height[low] > left_max){
				left_max = height[low]
			} else{
				cap += left_max - height[low]
			}
			low++
		}else{
			if(height[high] > right_max){
				right_max = height[high]
			}else{
				cap += right_max - height[high]
			}
			high--
		}
	}
	return cap
}

//By row
func trap2(height []int) int{
	var l int = len(height)
	var max_height int = 0
	for i := 0;i < l;i++{
		if height[i] > max_height{
			max_height = height[i]
		}
	}
	var res int = 0
	for h := 1;h <= max_height;h++{
		left := 0
		right := l - 1
		for left < right && height[left] < h{
			left++
		}
		for right > left && height[right] < h{
			right--
		}
		for j := left;j <= right;j++{
			if height[j] < h{
				res++
			}
		}
	}
	return res
}

//By columns
func trap3(height []int) int{
	return 0
}