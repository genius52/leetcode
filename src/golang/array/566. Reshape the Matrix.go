package array

//Input:
//nums =
//[[1,2],
// [3,4]]
//r = 1, c = 4
//Output:
//[[1,2,3,4]]
//Explanation:
//The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix,
//fill it row by row by using the previous list.
func MatrixReshape(nums [][]int, r int, c int) [][]int {
	var rows int = len(nums)
	var columns int = len(nums[0])
	if rows * columns != r * c{
		return nums
	}
	var res [][]int = make([][]int,r)
	for i := 0;i < r;i++{
		res[i] = make([]int,c)
	}
	var index int = 0
	for i := 0;i < r;i++{
		for j := 0;j < c;j++{
			pre_r := index / columns
			pre_c := index % columns
			res[i][j] = nums[pre_r][pre_c]
			index++
		}
	}
	return res
}