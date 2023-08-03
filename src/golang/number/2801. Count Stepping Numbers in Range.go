package number

func dp_countSteppingNumbers(cur string, low string, l1 int, high string, l2 int) int {
	if len(cur) > l2 {
		return 0
	}

}

func countSteppingNumbers(low string, high string) int {
	var l1 int = len(low)
	var l2 int = len(high)
	//var dp [][10]int = make([][10]int, l2)
	return dp_countSteppingNumbers("", low, l1, high, l2)
}
