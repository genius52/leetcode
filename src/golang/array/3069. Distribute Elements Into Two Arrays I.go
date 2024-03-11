package array

func ResultArray(nums []int) []int {
	var arr1 []int = []int{nums[0]}
	var arr2 []int = []int{nums[1]}
	var idx1 int = 0
	var idx2 int = 0
	for i := 2; i < len(nums); i++ {
		if arr1[idx1] > arr2[idx2] {
			arr1 = append(arr1, nums[i])
			idx1++
		} else {
			arr2 = append(arr2, nums[i])
			idx2++
		}
	}
	return append(arr1, arr2...)
}
