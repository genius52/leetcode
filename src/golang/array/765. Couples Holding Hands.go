package array

//Input: row = [0, 2, 1, 3]
//Output: 1
//Explanation: We only need to swap the second (row[1]) and third (row[2])
func minSwapsCouples(row []int) int {
	var l int = len(row)
	var res int = 0
	for i := 0;i < l;i += 2{
		if (row[i] ^ row[i + 1]) == 1{
			continue
		}
		res++
		for j := i + 1;j < l;j++{
			if row[j] == (row[i] ^ 1){
				row[i + 1],row[j] = row[j],row[i+1]
				break
			}
		}
	}
	return res
}