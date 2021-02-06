package array

//Here is my idea: instead of calculating area by height*width, we can think it in a cumulative way. In other words, sum water amount of each bin(width=1).
//Search from left to right and maintain a max height of left and right separately, which is like a one-side wall of partial container.
// Fix the higher one and flow water from the lower part. For example, if current height of left is lower, we fill water in the left bin.
// Until left meets right, we filled the whole container
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