package number

//You are given two non-increasing 0-indexed integer arrays nums1​​​​​​ and nums2​​​​​​.
//
//A pair of indices (i, j), where 0 <= i < nums1.length and 0 <= j < nums2.length,
//is valid if both i <= j and nums1[i] <= nums2[j]. The distance of the pair is j - i​​​​.
//
//Return the maximum distance of any valid pair (i, j). If there are no valid pairs, return 0.
//
//An array arr is non-increasing if arr[i-1] >= arr[i] for every 1 <= i < arr.length.

//Example 1:
//
//Input: nums1 = [55,30,5,4,2], nums2 = [100,20,10,10,5]
//Output: 2
//Explanation: The valid pairs are (0,0), (2,2), (2,3), (2,4), (3,3), (3,4), and (4,4).
//The maximum distance is 2 with pair (2,4).

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
