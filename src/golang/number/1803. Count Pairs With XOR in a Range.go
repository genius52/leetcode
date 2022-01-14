package number

type Binary_tie struct {
	cnt  int            //小于等于该节点的数量
	near [2]*Binary_tie //point to zero and one
}

func countPairs(nums []int, low int, high int) int {
	var root Binary_tie
	var l int = len(nums)
	for i := 0; i < l; i++ {
		var node *Binary_tie = &root
		for j := 31; j >= 0; j-- {
			var idx int = 0
			if nums[i]&1<<j == 1 {
				idx = 1
			}
			if node.near[idx] == nil {
				node.near[idx] = new(Binary_tie)
			}
			node.cnt++
			node = node.near[idx]
		}
	}
	return 0
}
