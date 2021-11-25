package number

import "sort"

func StoneGameVI(aliceValues []int, bobValues []int) int {
	var l int = len(aliceValues)
	var record [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		record[i] = []int{aliceValues[i],bobValues[i]}
	}
	sort.Slice(record, func(i, j int) bool {
		return (record[i][0] + record[i][1]) > (record[j][0] + record[j][1])
	})
	var alice int = 0
	var bob int = 0
	for i := 0;i < l;i++{
		if i | 1 == i{
			bob += record[i][1]
		}else{
			alice += record[i][0]
		}
	}
	if alice > bob{
		return 1
	}else if alice < bob{
		return -1
	}else{
		return 0
	}
}