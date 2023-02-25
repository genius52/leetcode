package number

func SquareFreeSubsets(nums []int) int {
	var res int = 0
	var allow []int = []int{1, 2, 3, 5, 6, 7, 10, 11, 13, 14, 15, 17, 19, 21, 22, 23, 26, 29, 30}
	var data map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(allow); j++ {
			if nums[i] == allow[j] {
				data[nums[i]]++
				break
			}
		}
	}
	//for _,_ := range data {
	//
	//}
	return res
}
