package array

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	var l1 int = len(arr1)
	var l2 int = len(arr2)
	var memo [][]int = make([][]int,l1)
	for i := 0;i < l1;i++{
		memo[i] = make([]int,l2)
	}
	return 0
}