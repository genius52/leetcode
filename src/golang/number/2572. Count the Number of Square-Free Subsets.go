package number

func SquareFreeSubsets(nums []int) int {
	var res int = 0
	var allow []int = []int{1, 2, 3, 5, 6, 7, 10, 11, 13, 14, 15, 17, 19, 21, 22, 23, 26, 29, 30}
	var record map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(allow); j++ {
			if nums[i] == allow[j] {
				record[nums[i]]++
				break
			}
		}
	}
	//var visited map[int]bool = make(map[int]bool)
	//for n1, cnt1 := range record {
	//	visited[n1] = true
	//	for n2, cnt2 := range record {
	//		if _, ok := visited[n2]; ok {
	//			continue
	//		}
	//		if gcd(n1, n2) == 1 {
	//
	//		}
	//	}
	//}
	return res
}
