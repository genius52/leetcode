package number

import "sort"

//Input: x = 2, y = 3, bound = 10
//Output: [2,3,4,5,7,9,10]
//Explanation:
//2 = 2^0 + 3^0
//3 = 2^1 + 3^0
//4 = 2^0 + 3^1
//5 = 2^1 + 3^1
//7 = 2^2 + 3^1
//9 = 2^3 + 3^0
//10 = 2^0 + 3^2
func PowerfulIntegers(x int, y int, bound int) []int {
	var record map[int]bool = make(map[int]bool)
	for i := 1;i <= bound;i = i * x{
		for j := 1;i + j <= bound;j = j * y{
			record[i + j] = true
			if(y == 1){
				break
			}
		}
		if(x == 1){
			break
		}
	}
	var keys []int
	for k := range record {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
