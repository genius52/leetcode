package number

//Input: nums1 = [55,30,5,4,2], nums2 = [100,20,10,10,5]
//Output: 2
//Explanation: The valid pairs are (0,0), (2,2), (2,3), (2,4), (3,3), (3,4), and (4,4).
//The maximum distance is 2 with pair (2,4).

//two pointer
func maxDistance(nums1 []int, nums2 []int) int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var index1 int = 0
	var index2 int = 0
	var res int = 0
	for index1 < l1 && index2 < l2 {
		if nums1[index1] <= nums2[index2] {
			var dis int = index2 - index1
			if dis >= res {
				res = dis
			}
			index2++
		} else {
			index1++
			if index1 > index2 {
				index2 = index1 + 1
			}
			rest_distance := l2 - 1 - index1
			if rest_distance <= res {
				break
			}
		}
	}
	return res
}

//binary search
func maxDistance2(nums1 []int, nums2 []int) int{
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var max_distance int = 0
	for i := 0;i < l1;i++{
		var left int = i
		var right int = l2 - 1
		var res int = i
		if l2 - 1 - i <= max_distance{
			break
		}
		for left <= right{
			mid := (left + right)/2
			if nums1[i] > nums2[mid]{
				right = mid - 1
				res = mid - 1
			}else{
				left = mid + 1
				res = mid
			}
		}
		if res >= 0 && res < l2 && nums1[i] <= nums2[res]{
			cur := res - i
			if cur > max_distance{
				max_distance = cur
			}
		}
	}
	return max_distance
}