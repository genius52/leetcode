package array

func findPeaks(mountain []int) []int {
	var res []int
	var l int = len(mountain)
	for i := 1; i < l-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			res = append(res, i)
		}
	}
	return res
}
