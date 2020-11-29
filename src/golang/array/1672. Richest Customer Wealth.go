package array

func maximumWealth(accounts [][]int) int {
	var res int = 0
	var rows int = len(accounts)
	var columns int = len(accounts[0])
	for i := 0;i < rows;i++{
		var money int = 0
		for j := 0;j < columns;j++{
			money += accounts[i][j]
		}
		if money > res{
			res = money
		}
	}
	return res
}