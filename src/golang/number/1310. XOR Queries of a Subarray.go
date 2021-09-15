package number

//Input: arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
//Output: [2,7,14,8]
func xorQueries(arr []int, queries [][]int) []int {
	var l int = len(arr)
	var l2 int = len(queries)
	var prefix []int = make([]int,l + 1)
	//prefix[0] = arr[0]
	for i := 0;i < l;i++{
		prefix[i + 1] = arr[i] ^ prefix[i]
	}
	var res []int = make([]int,l2)
	for i := 0;i < l2;i++{
		start := queries[i][0]
		end := queries[i][1]
		res[i] = prefix[end + 1] ^ prefix[start]
	}
	return res
}